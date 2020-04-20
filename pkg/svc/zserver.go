package svc

import (
	"context"
	"database/sql"
	albumpb "github.com/crud-grpc/pkg/pb"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type AlbumServiceServer struct{}

type Album struct {
	Title  string `json:"title"`
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
}

func NewAlbumServer() albumpb.AlbumServiceServer {
	return &AlbumServiceServer{}
}

func (a *AlbumServiceServer) GetAlbum(ctx context.Context, req *albumpb.Albumreq) (*albumpb.Albumresp, error) {

	db, err := sql.Open("mysql", "root:$pw@tcp(127.0.0.1:3306)/typicode")
	if err != nil {
		return nil, err
	}

	defer db.Close()
	result, err := db.Query("SELECT id, userId, title from album where id=?", req.GetId())
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var album Album
	for result.Next() {
		err := result.Scan(&album.Id, &album.UserId, &album.Title)
		if err != nil {
			return nil, err
		}
	}
	response := &albumpb.Albumresp{
		Album: &albumpb.Album{
			Id:     string(album.Id),
			UserId: string(album.UserId),
			Title:  album.Title,
		},
	}
	log.Infof("sending response back")
	return response, nil

}

func (a *AlbumServiceServer) ListAlbum(req *albumpb.ListAlbumRequest, stream albumpb.AlbumService_ListAlbumServer) error {

        db, err := sql.Open("mysql", "root:$pw@tcp(127.0.0.1:3306)/typicode")
        if err != nil {
                panic(err.Error())
        }

        defer db.Close()
        result, err := db.Query("SELECT id, userId, title from album")
        if err != nil {
               return err
        }
        defer result.Close()
        var album Album
        for result.Next() {
                err := result.Scan(&album.Id, &album.UserId, &album.Title)
                if err != nil {
               return err
                }
		stream.Send (&albumpb.ListAlbumResponse{
			 Album: &albumpb.Album{
                        Id:     string(album.Id),
                        UserId: string(album.UserId),
                        Title:  album.Title,
                },

		})
        }
	return nil
}
