# config

公用配置库


示例：

```
go get -u github.com/tristin2024/config

import "github.com/tristin2024/config"

if err := config.Viper.UnmarshalKey("postgresql", &PostgreSQL); err != nil {
		panic(err)
	}
```

