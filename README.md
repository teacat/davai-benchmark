# Davai Benchmark

此測試透過 `go run main.go` 進行，這會在最終產生一個 `benchmark.csv` 報表供檢視與方便生成圖表。

這份測試是基於下列電腦所執行的。

```
MacBook Air 13" (Mid-2013)
1.7 GHz Intel Core i7 (4650U)
8 GB 1600 MHz DDR3
Intel HD Graphics 5000 1536 MB
```

測試結果如下。

```
          TotalRequest   ReadSize     Reqs/s      Trans/s      Avg
Davai         21022       2.29MB     21583.24     2.35MB    231.661µs
Gramework     21364       3.97MB     22005.52     4.09MB    227.215µs
HTTPRouter    21100       2.29MB     21632.75     2.35MB    231.131µs
Martini       17949       1.95MB     18313.46     1.99MB    273.023µs
Pat           19397       2.11MB     19869.14     2.16MB    251.646µs
Gin           20718       2.25MB     21218.12     2.31MB    235.647µs
Mux           19123       2.08MB     19548.15     2.13MB    255.778µs
HTTPTreeMux   20847       2.27MB     21354.9      2.32MB    234.138µs
Echo          20508       2.23MB     21005.71     2.28MB    238.03µs
Beego         17619       2.32MB     17983.01     2.37MB    278.04µs
```