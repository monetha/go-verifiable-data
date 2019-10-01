package contracts

//go:generate abigen --abi ./code/IPassportLogic.abi --out IPassportLogic.go --pkg contracts --type IPassportLogicContract
//go:generate abigen --abi ./code/IPassportLogicRegistry.abi --out IPassportLogicRegistry.go --pkg contracts --type IPassportLogicRegistryContract
//go:generate abigen --abi ./code/Passport.abi --bin ./code/Passport.bin --out Passport.go --pkg contracts --type PassportContract
//go:generate abigen --abi ./code/PassportFactory.abi --bin ./code/PassportFactory.bin --out PassportFactory.go --pkg contracts --type PassportFactoryContract
//go:generate abigen --abi ./code/PassportLogic.abi --bin ./code/PassportLogic.bin --out PassportLogic.go --pkg contracts --type PassportLogicContract
//go:generate abigen --abi ./code/PassportLogicRegistry.abi --bin ./code/PassportLogicRegistry.bin --out PassportLogicRegistry.go --pkg contracts --type PassportLogicRegistryContract
//go:generate abigen --abi ./code/FactProviderRegistry.abi --bin ./code/FactProviderRegistry.bin --out FactProviderRegistry.go --pkg contracts --type FactProviderRegistryContract
