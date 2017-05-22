package service

import (
	"model"

	daoApi "api/dao_service"
	"encoding/json"

	"github.com/astaxie/beego"
)

type AlgorithmService struct {
}

func (this *AlgorithmService) Create(algorithm *model.Algorithm) (*model.Algorithm, error) {
	var err error
	var newAlgorithm *model.Algorithm

	// get admin
	beego.Debug("->get admin")
	var admin *model.User
	admin, err = daoApi.UserDaoApi.GetByName("admin")
	if err != nil {
		return nil, err
	}

	// get resource
	beego.Debug("->get admin resource")
	var adminResource *model.Resource
	adminResource, err = daoApi.ResourceDaoApi.GetByUserId(admin.Id)
	if err != nil {
		return nil, err
	}

	var algorithms []*model.Algorithm
	err = json.Unmarshal(([]byte)(adminResource.AlgorithmResource), &algorithms)
	if err != nil {
		return nil, err
	}

	// get by name and version
	beego.Debug("->get by name and version")
	var existed *model.Algorithm
	existed, err = daoApi.AlgorithmDaoApi.GetByNameAndVersion(algorithm.Name, algorithm.Version)

	//check if existed
	beego.Debug("->check if existed")
	if err != nil {
		// create
		beego.Debug("->create")
		algorithm.Id = 0
		algorithm.Stars = 0
		algorithm.Downloads = 0
		newAlgorithm, err = daoApi.AlgorithmDaoApi.Create(algorithm)
		if err != nil {
			return nil, err
		}
		algorithms = append(algorithms, newAlgorithm)
	} else {
		// update
		beego.Debug("->update")
		algorithm.Id = 0
		newAlgorithm, err = daoApi.AlgorithmDaoApi.Update(algorithm)
		if err != nil {
			return nil, err
		}

		for _, a := range algorithms {
			if a.Id == existed.Id {
				a = newAlgorithm
				break
			}
		}
	}

	// update admin resource
	beego.Debug("->update admin resource")
	var resultBody []byte
	resultBody, err = json.Marshal(&algorithms)
	if err != nil {
		beego.Debug("json Unmarshal data failed")
		return nil, err
	}
	adminResource.AlgorithmResource = string(resultBody)

	_, err = daoApi.ResourceDaoApi.Update(adminResource)
	if err != nil {
		return nil, err
	}

	return newAlgorithm, err
}

func (this *AlgorithmService) GetAll() ([]*model.Algorithm, error) {
	var err error
	var algorithms []*model.Algorithm

	algorithms, err = daoApi.AlgorithmDaoApi.GetAll()
	if err != nil {
		return nil, err
	}

	return algorithms, err
}

func (this *AlgorithmService) GetById(id int64) (*model.Algorithm, error) {
	var err error
	var algorithm *model.Algorithm

	algorithm, err = daoApi.AlgorithmDaoApi.GetById(id)
	if err != nil {
		return nil, err
	}

	return algorithm, err
}

func (this *AlgorithmService) Update(algorithm *model.Algorithm) (*model.Algorithm, error) {
	var err error
	var newAlgorithm *model.Algorithm

	newAlgorithm, err = daoApi.AlgorithmDaoApi.Update(algorithm)
	if err != nil {
		return nil, err
	}

	return newAlgorithm, err
}

func (this *AlgorithmService) DeleteById(id int64) error {
	var err error

	err = daoApi.AlgorithmDaoApi.DeleteById(id)
	if err != nil {
		return err
	}

	return err
}

func (this *AlgorithmService) Delete(algorithms []*model.Algorithm) error {
	var err error

	// delete algorithms
	beego.Debug("->delete algorithms")
	for _, a := range algorithms {
		err = daoApi.AlgorithmDaoApi.DeleteById(a.Id)
		if err != nil {
			return err
		}
	}

	// get admin
	beego.Debug("->get admin")
	var admin *model.User
	admin, err = daoApi.UserDaoApi.GetByName("admin")
	if err != nil {
		return err
	}

	// get resource
	beego.Debug("->get admin resource")
	var adminResource *model.Resource
	adminResource, err = daoApi.ResourceDaoApi.GetByUserId(admin.Id)
	if err != nil {
		return err
	}

	var adminAlgorithms []*model.Algorithm
	err = json.Unmarshal(([]byte)(adminResource.AlgorithmResource), &adminAlgorithms)
	if err != nil {
		return err
	}

	var finalAlgorithms []*model.Algorithm
	isDeleted := false
	for _, existed := range adminAlgorithms {
		isDeleted = false
		for _, deleted := range algorithms {
			if existed.Id == deleted.Id {
				isDeleted = true
				break
			}
		}

		if !isDeleted {
			finalAlgorithms = append(finalAlgorithms, existed)
		}
	}

	// update admin resource
	beego.Debug("->update admin resource")
	var resultBody []byte
	resultBody, err = json.Marshal(&algorithms)
	if err != nil {
		beego.Debug("json Unmarshal data failed")
		return err
	}
	adminResource.AlgorithmResource = string(resultBody)

	_, err = daoApi.ResourceDaoApi.Update(adminResource)
	if err != nil {
		return err
	}

	return err
}
