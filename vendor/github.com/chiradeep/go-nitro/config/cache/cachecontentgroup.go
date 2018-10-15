package cache

type Cachecontentgroup struct {
	Absexpiry              interface{} `json:"absexpiry,omitempty"`
	Absexpirygmt           interface{} `json:"absexpirygmt,omitempty"`
	Alwaysevalpolicies     string      `json:"alwaysevalpolicies,omitempty"`
	Builtin                interface{} `json:"builtin,omitempty"`
	Cache304hits           int         `json:"cache304hits,omitempty"`
	Cachecells             int         `json:"cachecells,omitempty"`
	Cachecontrol           string      `json:"cachecontrol,omitempty"`
	Cachegroupincarnation  int         `json:"cachegroupincarnation,omitempty"`
	Cachenon304hits        int         `json:"cachenon304hits,omitempty"`
	Cachenuminvalpolicy    int         `json:"cachenuminvalpolicy,omitempty"`
	Disklimit              int         `json:"disklimit,omitempty"`
	Expireatlastbyte       string      `json:"expireatlastbyte,omitempty"`
	Flags                  int         `json:"flags,omitempty"`
	Flashcache             string      `json:"flashcache,omitempty"`
	Heurexpiryparam        int         `json:"heurexpiryparam,omitempty"`
	Hitparams              interface{} `json:"hitparams,omitempty"`
	Hitselector            string      `json:"hitselector,omitempty"`
	Host                   string      `json:"host,omitempty"`
	Ignoreparamvaluecase   string      `json:"ignoreparamvaluecase,omitempty"`
	Ignorereloadreq        string      `json:"ignorereloadreq,omitempty"`
	Ignorereqcachinghdrs   string      `json:"ignorereqcachinghdrs,omitempty"`
	Insertage              string      `json:"insertage,omitempty"`
	Insertetag             string      `json:"insertetag,omitempty"`
	Insertvia              string      `json:"insertvia,omitempty"`
	Invalparams            interface{} `json:"invalparams,omitempty"`
	Invalrestrictedtohost  string      `json:"invalrestrictedtohost,omitempty"`
	Invalselector          string      `json:"invalselector,omitempty"`
	Lazydnsresolve         string      `json:"lazydnsresolve,omitempty"`
	Markercells            int         `json:"markercells,omitempty"`
	Matchcookies           string      `json:"matchcookies,omitempty"`
	Maxressize             int         `json:"maxressize,omitempty"`
	Memdusage              int         `json:"memdusage,omitempty"`
	Memlimit               int         `json:"memlimit,omitempty"`
	Memusage               int         `json:"memusage,omitempty"`
	Minhits                int         `json:"minhits,omitempty"`
	Minressize             int         `json:"minressize,omitempty"`
	Name                   string      `json:"name,omitempty"`
	Persist                string      `json:"persist,omitempty"`
	Persistha              string      `json:"persistha,omitempty"`
	Pinned                 string      `json:"pinned,omitempty"`
	Policyname             interface{} `json:"policyname,omitempty"`
	Polleverytime          string      `json:"polleverytime,omitempty"`
	Prefetch               string      `json:"prefetch,omitempty"`
	Prefetchcur            int         `json:"prefetchcur,omitempty"`
	Prefetchmaxpending     int         `json:"prefetchmaxpending,omitempty"`
	Prefetchperiod         int         `json:"prefetchperiod,omitempty"`
	Prefetchperiodmillisec int         `json:"prefetchperiodmillisec,omitempty"`
	Query                  string      `json:"query,omitempty"`
	Quickabortsize         int         `json:"quickabortsize,omitempty"`
	Relexpiry              int         `json:"relexpiry,omitempty"`
	Relexpirymillisec      int         `json:"relexpirymillisec,omitempty"`
	Removecookies          string      `json:"removecookies,omitempty"`
	Selectorvalue          string      `json:"selectorvalue,omitempty"`
	Tosecondary            string      `json:"tosecondary,omitempty"`
	Type                   string      `json:"type,omitempty"`
	Weaknegrelexpiry       int         `json:"weaknegrelexpiry,omitempty"`
	Weakposrelexpiry       int         `json:"weakposrelexpiry,omitempty"`
}
