package contracts

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// MedicineContract manages medicine lifecycle
type MedicineContract struct {
	contractapi.Contract
}

type Medicine struct {
	AssetType     string  `json:"assetType"`
	MedicineID    string  `json:"medicineId"`
	Name          string  `json:"name"`
	Manufacturer  string  `json:"manufacturer"`
	Price         float64 `json:"price"`
	Composition   string  `json:"composition"`
	ExpiryDate    string  `json:"expiryDate"`
	Status        string  `json:"status"`
}

// MedicineExists checks if medicine exists in the world state
func (c *MedicineContract) MedicineExists(ctx contractapi.TransactionContextInterface, medicineID string) (bool, error) {
	data, err := ctx.GetStub().GetState(medicineID)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return data != nil, nil
}

// CreateMedicine allows only manufacturer to create a medicine
func (c *MedicineContract) CreateMedicine(ctx contractapi.TransactionContextInterface, medicineID, name, manufacturer, composition, expiryDate string, price float64) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgID != "manufacturer-medicine-com" {
		return "", fmt.Errorf("unauthorized: %s cannot create medicine", clientOrgID)
	}

	exists, err := c.MedicineExists(ctx, medicineID)
	if err != nil {
		return "", err
	}
	if exists {
		return "", fmt.Errorf("medicine %s already exists", medicineID)
	}

	medicine := Medicine{
		AssetType:    "medicine",
		MedicineID:   medicineID,
		Name:         name,
		Manufacturer: manufacturer,
		Composition:  composition,
		ExpiryDate:   expiryDate,
		Price:        price,
		Status:       "Manufactured",
	}

	bytes, err := json.Marshal(medicine)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(medicineID, bytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("medicine %s created successfully", medicineID), nil
}

// ReadMedicine retrieves a medicine by ID
func (c *MedicineContract) ReadMedicine(ctx contractapi.TransactionContextInterface, medicineID string) (*Medicine, error) {
	data, err := ctx.GetStub().GetState(medicineID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if data == nil {
		return nil, fmt.Errorf("medicine %s does not exist", medicineID)
	}

	var medicine Medicine
	err = json.Unmarshal(data, &medicine)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}

	return &medicine, nil
}

// UpdatePrice allows only dealer to update medicine price
func (c *MedicineContract) UpdatePrice(ctx contractapi.TransactionContextInterface, medicineID string, newPrice float64) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgID != "dealer-medicine-com" {
		return "", fmt.Errorf("unauthorized: %s cannot update price", clientOrgID)
	}

	medicineBytes, err := ctx.GetStub().GetState(medicineID)
	if err != nil {
		return "", err
	}
	if medicineBytes == nil {
		return "", fmt.Errorf("medicine %s does not exist", medicineID)
	}

	var medicine Medicine
	err = json.Unmarshal(medicineBytes, &medicine)
	if err != nil {
		return "", err
	}

	medicine.Price = newPrice

	updatedBytes, err := json.Marshal(medicine)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(medicineID, updatedBytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("price for medicine %s updated to %.2f", medicineID, newPrice), nil
}
