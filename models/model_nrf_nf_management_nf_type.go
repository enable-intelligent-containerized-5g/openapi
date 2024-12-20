package models

type NrfNfManagementNfType string

// List of NrfNFManagementNFType
const (
	NrfNfManagementNfType_NRF        NrfNfManagementNfType = "NRF"
	NrfNfManagementNfType_UDM        NrfNfManagementNfType = "UDM"
	NrfNfManagementNfType_AMF        NrfNfManagementNfType = "AMF"
	NrfNfManagementNfType_SMF        NrfNfManagementNfType = "SMF"
	NrfNfManagementNfType_AUSF       NrfNfManagementNfType = "AUSF"
	NrfNfManagementNfType_NEF        NrfNfManagementNfType = "NEF"
	NrfNfManagementNfType_PCF        NrfNfManagementNfType = "PCF"
	NrfNfManagementNfType_SMSF       NrfNfManagementNfType = "SMSF"
	NrfNfManagementNfType_NSSF       NrfNfManagementNfType = "NSSF"
	NrfNfManagementNfType_UDR        NrfNfManagementNfType = "UDR"
	NrfNfManagementNfType_LMF        NrfNfManagementNfType = "LMF"
	NrfNfManagementNfType_GMLC       NrfNfManagementNfType = "GMLC"
	NrfNfManagementNfType__5_G_EIR   NrfNfManagementNfType = "5G_EIR"
	NrfNfManagementNfType_SEPP       NrfNfManagementNfType = "SEPP"
	NrfNfManagementNfType_UPF        NrfNfManagementNfType = "UPF"
	NrfNfManagementNfType_N3_IWF     NrfNfManagementNfType = "N3IWF"
	NrfNfManagementNfType_AF         NrfNfManagementNfType = "AF"
	NrfNfManagementNfType_UDSF       NrfNfManagementNfType = "UDSF"
	NrfNfManagementNfType_BSF        NrfNfManagementNfType = "BSF"
	NrfNfManagementNfType_CHF        NrfNfManagementNfType = "CHF"
	NrfNfManagementNfType_NWDAF      NrfNfManagementNfType = "NWDAF"
	NrfNfManagementNfType_PCSCF      NrfNfManagementNfType = "PCSCF"
	NrfNfManagementNfType_CBCF       NrfNfManagementNfType = "CBCF"
	NrfNfManagementNfType_HSS        NrfNfManagementNfType = "HSS"
	NrfNfManagementNfType_UCMF       NrfNfManagementNfType = "UCMF"
	NrfNfManagementNfType_SOR_AF     NrfNfManagementNfType = "SOR_AF"
	NrfNfManagementNfType_SPAF       NrfNfManagementNfType = "SPAF"
	NrfNfManagementNfType_MME        NrfNfManagementNfType = "MME"
	NrfNfManagementNfType_SCSAS      NrfNfManagementNfType = "SCSAS"
	NrfNfManagementNfType_SCEF       NrfNfManagementNfType = "SCEF"
	NrfNfManagementNfType_SCP        NrfNfManagementNfType = "SCP"
	NrfNfManagementNfType_NSSAAF     NrfNfManagementNfType = "NSSAAF"
	NrfNfManagementNfType_ICSCF      NrfNfManagementNfType = "ICSCF"
	NrfNfManagementNfType_SCSCF      NrfNfManagementNfType = "SCSCF"
	NrfNfManagementNfType_DRA        NrfNfManagementNfType = "DRA"
	NrfNfManagementNfType_IMS_AS     NrfNfManagementNfType = "IMS_AS"
	NrfNfManagementNfType_AANF       NrfNfManagementNfType = "AANF"
	NrfNfManagementNfType__5_G_DDNMF NrfNfManagementNfType = "5G_DDNMF"
	NrfNfManagementNfType_NSACF      NrfNfManagementNfType = "NSACF"
	NrfNfManagementNfType_MFAF       NrfNfManagementNfType = "MFAF"
	NrfNfManagementNfType_EASDF      NrfNfManagementNfType = "EASDF"
	NrfNfManagementNfType_DCCF       NrfNfManagementNfType = "DCCF"
	NrfNfManagementNfType_MB_SMF     NrfNfManagementNfType = "MB_SMF"
	NrfNfManagementNfType_TSCTSF     NrfNfManagementNfType = "TSCTSF"
	NrfNfManagementNfType_ADRF       NrfNfManagementNfType = "ADRF"
	NrfNfManagementNfType_GBA_BSF    NrfNfManagementNfType = "GBA_BSF"
	NrfNfManagementNfType_CEF        NrfNfManagementNfType = "CEF"
	NrfNfManagementNfType_MB_UPF     NrfNfManagementNfType = "MB_UPF"
	NrfNfManagementNfType_NSWOF      NrfNfManagementNfType = "NSWOF"
	NrfNfManagementNfType_PKMF       NrfNfManagementNfType = "PKMF"
	NrfNfManagementNfType_MNPF       NrfNfManagementNfType = "MNPF"
	NrfNfManagementNfType_SMS_GMSC   NrfNfManagementNfType = "SMS_GMSC"
	NrfNfManagementNfType_SMS_IWMSC  NrfNfManagementNfType = "SMS_IWMSC"
	NrfNfManagementNfType_MBSF       NrfNfManagementNfType = "MBSF"
	NrfNfManagementNfType_MBSTF      NrfNfManagementNfType = "MBSTF"
	NrfNfManagementNfType_PANF       NrfNfManagementNfType = "PANF"
)
