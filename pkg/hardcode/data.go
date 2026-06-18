package hardcode

// HardCodeCategory 硬编码规则分类
type HardCodeCategory struct {
	Name     string
	Patterns []string
}

// HardCodeRules 按类别定义硬编码规则
var HardCodeRules = []HardCodeCategory{
	{
		Name: "HAE敏感信息",
		Patterns: []string{
			`(?i)password\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)passwd\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)pwd\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)username\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)secret\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)key\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)token\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)auth\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)pass\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)login\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)email\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)account\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)access_key\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)secret_key\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)client_secret\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)client_id\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)api[_-]?key\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)api[_-]?secret\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)access[_-]?token\s*=\s*['"][^'"]{3,}['"]`,
			`(?i)refresh[_-]?token\s*=\s*['"][^'"]{3,}['"]`,
		},
	},
	{
		Name: "私钥和证书",
		Patterns: []string{
			"-----BEGIN DSA PRIVATE KEY-----",
			"-----BEGIN EC PRIVATE KEY-----",
			"-----BEGIN PGP PRIVATE KEY BLOCK-----",
			"-----BEGIN RSA PRIVATE KEY-----",
		},
	},
	{
		Name: "API密钥和令牌",
		Patterns: []string{
			`[aA][pP][iI]_?[kK][eE][yY].{0,20}['|\"][0-9a-zA-Z]{32,45}['|\"]`,
			`api_key=[0-9a-zA-Z]+`,
			`key-[0-9a-zA-Z]{32}`,
			`AIza[0-9A-Za-z\-_]{35}`,
		},
	},
	{
		Name: "OAuth和认证令牌",
		Patterns: []string{
			`access_token=[0-9a-zA-Z]+`,
			`access_token\$production\$[0-9a-z]{16}\$[0-9a-f]{32}`,
			`ya29\.[0-9A-Za-z\-_]+`,
			`eyJhbGciOiJ`,
			`EAACEdEose0cBA[0-9A-Za-z]+`,
		},
	},
	{
		Name: "云平台凭证",
		Patterns: []string{
			`((?:A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16})`,
			`amzn\.mws\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`,
			`[hH][eE][rR][oO][kK][uU].{0,20}[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}`,
		},
	},
	{
		Name: "服务账号凭证",
		Patterns: []string{
			`"service_account\"`,
			`[0-9]+-[0-9A-Za-z_]{32}\.apps\.googleusercontent\.com`,
			`[0-9]+:AA[0-9A-Za-z\-_]{33}`,
		},
	},
	{
		Name: "支付相关密钥",
		Patterns: []string{
			`rk_live_[0-9a-zA-Z]{24}`,
			`sk_live_[0-9a-z]{32}`,
			`sk_live_[0-9a-zA-Z]{24}`,
			`sq0atp-[0-9A-Za-z\-_]{22}`,
			`sq0csp-[0-9A-Za-z\-_]{43}`,
		},
	},
	{
		Name: "平台服务密钥",
		Patterns: []string{
			`[fF][aA][cC][eE][bB][oO][oO][kK].{0,20}['|\"][0-9a-f]{32}['|\"]`,
			`[gG][iI][tT][hH][uU][bB].{0,20}['|\"][0-9a-zA-Z]{35,40}['|\"]`,
			`https://hooks\.slack\.com/services/T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}`,
			`https:\/\/[a-zA-Z0-9]{40}@github\.com`,
		},
	},
	{
		Name: "其他敏感信息",
		Patterns: []string{
			`[sS][eE][cC][rR][eE][tT].{0,20}['|\"][0-9a-zA-Z]{32,45}['|\"]`,
			`[a-zA-Z]{3,10}://[^/\s:@]{3,20}:[^/\s:@]{3,20}@.{1,100}["'\s]`,
			`[0-9a-f]{32}-us[0-9]{1,2}`,
			`da2-[a-z0-9]{26}`,
			`SK[0-9a-fA-F]{32}`,
		},
	},
}
