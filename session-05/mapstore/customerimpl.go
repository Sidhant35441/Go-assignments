package mapstore

import (
	"assignments/session-05/domain"
)

type MapStore struct {
	store map[string]domain.Customer
}

func Create(k string, v domain.Customer) error {

}
