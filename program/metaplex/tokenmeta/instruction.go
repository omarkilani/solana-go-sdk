package tokenmeta

import (
	"github.com/near/borsh-go"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

type Instruction uint8

const (
	InstructionCreateMetadataAccount Instruction = iota
	InstructionUpdateMetadataAccunnt
	InstructionDeprecatedCreateMasterEdition
	InstructionDeprecatedMintNewEditionFromMasterEditionViaPrintingToken
	InstructionUpdatePrimarySaleHappenedViaToken
	InstructionDeprecatedSetReservationList
	InstructionDeprecatedCreateReservationList
	InstructionSignMetadata
	InstructionDeprecatedMintPrintingTokensViaToken
	InstructionDeprecatedMintPrintingTokens
	InstructionCreateMasterEdition
	InstructionMintNewEditionFromMasterEditionViaToken
	InstructionConvertMasterEditionV1ToV2
	InstructionMintNewEditionFromMasterEditionViaVaultProxy
	InstructionPuffMetadata
)

type CreateMetadataAccountParam struct {
	Metadata                common.PublicKey
	Mint                    common.PublicKey
	MintAuthority           common.PublicKey
	Payer                   common.PublicKey
	UpdateAuthority         common.PublicKey
	UpdateAuthorityIsSigner bool
	IsMutable               bool
	MintData                Data
}

func CreateMetadataAccount(param CreateMetadataAccountParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Data        Data
		IsMutable   bool
	}{
		Instruction: InstructionCreateMetadataAccount,
		Data:        param.MintData,
		IsMutable:   param.IsMutable,
	})

	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.MintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   param.UpdateAuthorityIsSigner,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type CreateMasterEditionParam struct {
	Edition         common.PublicKey
	Mint            common.PublicKey
	UpdateAuthority common.PublicKey
	MintAuthority   common.PublicKey
	Metadata        common.PublicKey
	Payer           common.PublicKey
	MaxSupply       *uint64
}

func CreateMasterEdition(param CreateMasterEditionParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		MaxSupply   *uint64
	}{
		Instruction: InstructionCreateMasterEdition,
		MaxSupply:   param.MaxSupply,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Edition,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.MintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.TokenProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type MintNewEditionFromMasterEditionViaTokeParam struct {
	NewMetaData                common.PublicKey
	NewEdition                 common.PublicKey
	MasterEdition              common.PublicKey
	NewMint                    common.PublicKey
	EditionMark                common.PublicKey
	NewMintAuthority           common.PublicKey
	Payer                      common.PublicKey
	TokenAccountOwner          common.PublicKey
	TokenAccount               common.PublicKey
	NewMetadataUpdateAuthority common.PublicKey
	MasterMetadata             common.PublicKey
	Edition                    uint64
}

func MintNewEditionFromMasterEditionViaToken(param MintNewEditionFromMasterEditionViaTokeParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Edition     uint64
	}{
		Instruction: InstructionMintNewEditionFromMasterEditionViaToken,
		Edition:     param.Edition,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.NewMetaData,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.NewEdition,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.MasterEdition,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.NewMint,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.EditionMark,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.NewMintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.TokenAccountOwner,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.TokenAccount,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.NewMetadataUpdateAuthority,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.MasterMetadata,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.TokenProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}
