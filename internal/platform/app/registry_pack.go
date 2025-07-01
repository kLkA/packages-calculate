package app

import (
	"homework/internal/pack"
	_packHTTP "homework/internal/pack/delivery/http"
	_packUsecase "homework/internal/pack/usecase"
)

func (r *Registry) GetPackServer() (*_packHTTP.PackServer, error) {
	if r.packServer == nil {
		packService, err := r.GetPackService()
		if err != nil {
			r.errors = append(r.errors, err)
			return nil, err
		}

		r.packServer = _packHTTP.NewServer(r.GetRouter(), packService)
	}
	return r.packServer, nil
}

func (r *Registry) GetPackService() (pack.Service, error) {
	if r.packService == nil {
		packService := _packUsecase.NewPackService()
		r.packService = packService
	}
	return r.packService, nil
}
