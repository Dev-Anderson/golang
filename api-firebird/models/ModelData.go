package models

import (
	"database/sql"

	_struct "api-firebird/struct"
)

type ModelGetData struct {
	DB *sql.DB
}

func (model ModelGetData) GetNotasAprovadas() (getStruct []_struct.StructData, err error) {
	row, err := model.DB.Query("SELECT 	IDE.NUMERO, 	IDE.SERIE, 	IDE.DATAEMISSAO, 	IDE.HORAEMISSAO, 	NFES.ID, 	NFES.NUMEROPROTOCOLO, 	TOT.VNF FROM NFE_TBNFES NFES INNER JOIN NFE_TBNFESIDE IDE ON (NFES.GID = IDE.GID) INNER JOIN NFE_TBCANCELAMENTONFE CAN ON (NFES.GID = CAN.GID) INNER JOIN NFE_TBNFESICMSTOT TOT ON (NFES.GID = TOT.NFE ) WHERE NFES.NUMEROPROTOCOLO IS NOT NULL AND CAN.CANCELADOEM  IS NULL")
	if err != nil {
		return nil, err
	} else {
		var _isiStruct []_struct.StructData
		var data _struct.StructData
		for row.Next() {
			err2 := row.Scan(
				&data.Numero,
				&data.Serie,
				&data.DataEmissao,
				&data.HoraEmissao,
				&data.Id,
				&data.Protocolo,
				&data.ValorNotafiscal,
			)
			if err2 != nil {
				return nil, err2
			} else {
				_data := _struct.StructData{
					Numero:          data.Numero,
					Serie:           data.Serie,
					DataEmissao:     data.DataEmissao,
					HoraEmissao:     data.HoraEmissao,
					Id:              data.Id,
					Protocolo:       data.Protocolo,
					ValorNotafiscal: data.ValorNotafiscal,
				}
				_isiStruct = append(_isiStruct, _data)
			}
		}
		return _isiStruct, nil
	}
}

func (model ModelGetData) GetNotasCanceladas() (getStruct []_struct.StructDataCanceladas, err error) {
	row, err := model.DB.Query("SELECT    NFES.GID,    IDE.NUMERO,    IDE.SERIE,    NFES.ID,    IDE.DATAEMISSAO,    IDE.HORAEMISSAO,    CAN.CANCELADOEM FROM NFE_TBNFES NFES INNER JOIN NFE_TBNFESIDE IDE ON NFES.GID = IDE.GID INNER JOIN NFE_TBCANCELAMENTONFE CAN ON NFES.GID = CAN.GID WHERE CAN.CANCELADOEM IS NOT NULL")
	if err != nil {
		return nil, err
	} else {
		var _isiStruct []_struct.StructDataCanceladas
		var data _struct.StructDataCanceladas
		for row.Next() {
			err2 := row.Scan(
				&data.Gid,
				&data.Numero,
				&data.Serie,
				&data.Id,
				&data.DataEmissao,
				&data.HoraEmissao,
				&data.CanceladoEm,
			)
			if err2 != nil {
				return nil, err2
			} else {
				_data := _struct.StructDataCanceladas{
					Gid:         data.Gid,
					Numero:      data.Numero,
					Serie:       data.Serie,
					Id:          data.Id,
					DataEmissao: data.DataEmissao,
					HoraEmissao: data.HoraEmissao,
					CanceladoEm: data.CanceladoEm,
				}
				_isiStruct = append(_isiStruct, _data)
			}
		}
		return _isiStruct, nil
	}
}

func (model ModelGetData) GetNfeAprovada() (getStruct []_struct.StructDataNfe, err error) {
	row, err := model.DB.Query("SELECT 	IDE.NUMERO, 	IDE.SERIE, 	IDE.DATAEMISSAO, 	IDE.HORAEMISSAO, 	NFES.ID, 	NFES.NUMEROPROTOCOLO, 	TOT.VNF FROM NFE_TBNFES NFES INNER JOIN NFE_TBNFESIDE IDE ON (NFES.GID = IDE.GID) INNER JOIN NFE_TBCANCELAMENTONFE CAN ON (NFES.GID = CAN.GID) INNER JOIN NFE_TBNFESICMSTOT TOT ON (NFES.GID = TOT.NFE ) WHERE NFES.NUMEROPROTOCOLO IS NOT NULL AND CAN.CANCELADOEM  IS NULL AND IDE.MOD  = 55")
	if err != nil {
		return nil, err
	} else {
		var _isiStruct []_struct.StructDataNfe
		var data _struct.StructDataNfe
		for row.Next() {
			err2 := row.Scan(
				&data.Gid,
				&data.Numero,
				&data.Serie,
				&data.DataEmissao,
				&data.HoraEmissao,
				&data.Id,
				&data.Protocolo,
			)
			if err2 != nil {
				return nil, err2
			} else {
				_data := _struct.StructDataNfe{
					Gid:         data.Gid,
					Numero:      data.Numero,
					Serie:       data.Serie,
					Id:          data.Id,
					DataEmissao: data.DataEmissao,
					HoraEmissao: data.HoraEmissao,
					Protocolo:   data.Protocolo,
				}
				_isiStruct = append(_isiStruct, _data)
			}
		}
		return _isiStruct, nil
	}
}
