package dao

import (
	"fmt"
	"github.com/jiuzi/sshm/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type MachineDao interface {
	Add(m *model.Machine) error
	Delete(id int) error
	SelectById(id int) (*model.Machine, error)
	SelectByName(name string) (*model.Machine, error)
	SelectAll() ([]model.Machine, error)
	SelectLikeName(name string) ([]model.Machine, error)
	UpdateMachineById(machine *model.Machine) error
}

type machineDaoImpl struct {
	Db *gorm.DB
}

func NewMachineDao(db *gorm.DB) *machineDaoImpl {
	return &machineDaoImpl{Db: db}
}

func (m machineDaoImpl) Add(machine *model.Machine) error {
	return errors.WithStack(m.Db.Create(machine).Error)
}

func (m machineDaoImpl) Delete(id int) error {
	return errors.WithStack(m.Db.Delete(&model.Machine{}, id).Error)
}

func (m machineDaoImpl) SelectById(id int) (*model.Machine, error) {
	var machine model.Machine
	machine.ID = uint(id)
	if err := m.Db.First(&machine).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &machine, nil

}

func (m machineDaoImpl) SelectByName(name string) (*model.Machine, error) {
	var machine model.Machine
	if err := m.Db.First(&machine, "name=?", name).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &machine, nil
}

func (m machineDaoImpl) SelectAll() ([]model.Machine, error) {
	/*	db := dao.DB
		var apsRoutes []ApsRoute
			result := db.Where("yn=1").Find(&apsRoutes)*/
	var machines []model.Machine
	err := m.Db.Find(&machines).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return machines, nil
}

func (m machineDaoImpl) SelectLikeName(arg string) ([]model.Machine, error) {
	var machines []model.Machine
	likeArg := fmt.Sprintf("%%%s%%", arg)
	err := m.Db.Where("name LIKE ?", likeArg).Or("Host LIKE ?", likeArg).Find(&machines).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return machines, nil
}

func (m machineDaoImpl) UpdateMachineById(machine *model.Machine) error {
	if machine == nil || machine.ID == 0 {
		return errors.New("ID must be specified!")
	}
	return errors.WithStack(
		m.Db.Model(machine).Updates(&model.Machine{
			Name:     machine.Name,
			Host:     machine.Host,
			Ip:       machine.Ip,
			Port:     machine.Port,
			User:     machine.User,
			Password: machine.Password,
			Key:      machine.Key,
			Type:     machine.Type,
		}).Error,
	)

}
