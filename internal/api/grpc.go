package api

import (
	"context"
	"fmt"

	"github.com/holin20/catcatcat/internal/ent/schema"
	"github.com/holin20/catcatcat/pkg/ezgo"
	"github.com/holin20/catcatcat/pkg/ezgo/orm"
	catcatcatpb "github.com/holin20/catcatcat/proto/catcatcat"
)

type grpcServer struct {
	//catcatcatpb.UnimplementedCatcatcatServer

	db *ezgo.PostgresDB
}

func NewGrpcServer() *grpcServer {
	db, err := ezgo.NewLocalPostgresDB("postgres", "postgres", 54320, "postgres")
	ezgo.AssertNoError(err, "failed to NewLocalPostgresDB")

	return &grpcServer{
		db: db,
	}
}

func (s *grpcServer) ListCats(
	ctx context.Context,
	in *catcatcatpb.ListCatsRequest,
) (*catcatcatpb.ListCatsResponse, error) {
	catById, err := orm.Load(s.db, schema.CatSchema)
	if ezgo.IsErr(err) {
		return nil, ezgo.NewCause(err, "LoadCats")
	}

	var catsProto []*catcatcatpb.Cat
	for _, cat := range catById {
		catsProto = append(catsProto, &catcatcatpb.Cat{
			CatId: cat.CatId,
			Name:  cat.Name,
		})
	}

	return &catcatcatpb.ListCatsResponse{
		Cats: catsProto,
	}, nil
}

func (s grpcServer) GetCdps(
	ctx context.Context,
	req *catcatcatpb.GetCdpsRequest,
) (*catcatcatpb.GetCdpsResponse, error) {
	fmt.Println(req.LastN)
	cdpPacks, err := orm.LoadLastN(s.db, schema.CdpSchema, &schema.Cdp{CatId: req.CatId}, int(req.LastN))
	if ezgo.IsErr(err) {
		return nil, ezgo.NewCause(err, "LoadLastN for cdp")
	}

	var cdpsProto []*catcatcatpb.Cdp
	for _, cdpPack := range cdpPacks {
		ts, cdp := cdpPack.Unpack()
		cdpsProto = append(cdpsProto, &catcatcatpb.Cdp{
			Ts:      ts,
			Price:   cdp.Price,
			InStock: cdp.InStock,
		})
	}

	return &catcatcatpb.GetCdpsResponse{
		Cat:  &catcatcatpb.Cat{CatId: req.CatId},
		Cdps: cdpsProto,
	}, nil
}
