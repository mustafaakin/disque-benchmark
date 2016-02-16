# disque-benchmark

Simple disque benchmark written in Go. Some errors exist right now


```txt
$ go run bench.go 
Msg:    1 N:   10000 Threads:    1 
[SET]  301.55844ms    10000 33161.068/s
[GET] 329.602567ms    10000 30339.570/s
======================================
Msg:    4 N:   10000 Threads:    1 
[SET] 367.631629ms    10000 27201.142/s
[GET] 308.445449ms    10000 32420.644/s
======================================
Msg:    8 N:   10000 Threads:    1 
[SET] 366.606958ms    10000 27277.169/s
[GET] 383.547668ms    10000 26072.379/s
======================================
Msg:   16 N:   10000 Threads:    1 
[SET] 379.026831ms    10000 26383.356/s
[GET] 315.885386ms    10000 31657.052/s
======================================
Msg:   20 N:   10000 Threads:    1 
[SET] 391.415264ms    10000 25548.314/s
[GET] 312.883835ms    10000 31960.744/s
======================================
Msg:  200 N:   10000 Threads:    1 
[SET] 359.679357ms    10000 27802.541/s
[GET] 344.532667ms    10000 29024.824/s
======================================
Msg: 1000 N:   10000 Threads:    1 
[SET]   449.5947ms    10000 22242.255/s
[GET] 337.849858ms    10000 29598.947/s
======================================
Msg:    1 N:   10000 Threads:    2 
[SET] 313.952211ms    20000 63703.963/s
[GET] 299.677076ms    20000 66738.505/s
======================================
Msg:    4 N:   10000 Threads:    2 
[SET] 323.625489ms    20000 61799.829/s
[GET] 362.583638ms    20000 55159.687/s
======================================
Msg:    8 N:   10000 Threads:    2 
[SET] 312.484287ms    20000 64003.218/s
[GET] 319.845494ms    20000 62530.192/s
=====================================
Msg:   16 N:   10000 Threads:    2 
[SET] 336.005432ms    20000 59522.847/s
[GET] 335.444334ms    20000 59622.411/s
======================================
Msg:   20 N:   10000 Threads:    2 
[SET] 325.865078ms    20000 61375.095/s
[GET] 302.141777ms    20000 66194.090/s
======================================
Msg:  200 N:   10000 Threads:    2 
[SET] 323.842953ms    20000 61758.330/s
[GET] 360.153765ms    20000 55531.837/s
======================================
Msg: 1000 N:   10000 Threads:    2 
[SET]  559.11169ms    20000 35771.028/s
[GET] 330.497001ms    20000 60514.921/s
======================================
Msg:    1 N:   10000 Threads:    4 
[SET] 533.431432ms    40000 74986.207/s
[GET] 520.301011ms    40000 76878.574/s
======================================
Msg:    4 N:   10000 Threads:    4 
[SET] 519.863463ms    40000 76943.280/s
[GET] 499.060259ms    40000 80150.642/s
======================================
Msg:    8 N:   10000 Threads:    4 
[SET] 543.767049ms    40000 73560.912/s
[GET] 573.214539ms    40000 69781.901/s
======================================
Msg:   16 N:   10000 Threads:    4 
[SET] 464.353643ms    40000 86141.243/s
[GET] 529.540811ms    40000 75537.143/s
======================================
Msg:   20 N:   10000 Threads:    4 
[SET] 591.186447ms    40000 67660.550/s
[GET] 551.684158ms    40000 72505.254/s
======================================
Msg:  200 N:   10000 Threads:    4 
[SET] 617.160594ms    40000 64812.952/s
[GET] 568.395662ms    40000 70373.514/s
======================================
Msg: 1000 N:   10000 Threads:    4 
[SET] 645.017041ms    40000 62013.865/s
[GET] 644.299604ms    40000 62082.919/s
======================================
Msg:    1 N:   10000 Threads:    8 
[SET] 1.042786749s    80000 76717.507/s
[GET] 997.140092ms    80000 80229.449/s
======================================
Msg:    4 N:   10000 Threads:    8 
[SET] 998.781417ms    80000 80097.606/s
[GET] 944.776492ms    80000 84676.112/s
======================================
Msg:    8 N:   10000 Threads:    8 
[SET]  902.98842ms    80000 88594.713/s
[GET] 1.030932405s    80000 77599.656/s
======================================
Msg:   16 N:   10000 Threads:    8 
[SET] 989.400443ms    80000 80857.049/s
[GET] 1.014081069s    80000 78889.156/s
======================================
Msg:   20 N:   10000 Threads:    8 
[SET] 884.055762ms    80000 90492.029/s
[GET] 994.324789ms    80000 80456.608/s
======================================
Msg:  200 N:   10000 Threads:    8 
[SET] 996.496589ms    80000 80281.258/s
[GET] 1.064933739s    80000 75122.045/s
======================================
Msg: 1000 N:   10000 Threads:    8 
[SET] 1.076843475s    80000 74291.206/s
[GET] 1.087490219s    80000 73563.880/s
======================================
Msg:    1 N:   10000 Threads:   10 
[SET] 1.161267138s   100000 86112.830/s
[GET] 1.184532903s   100000 84421.462/s
======================================
Msg:    4 N:   10000 Threads:   10 
[SET] 1.245156604s   100000 80311.183/s
[GET] 1.175385901s   100000 85078.441/s
======================================
Msg:    8 N:   10000 Threads:   10 
[SET] 1.096546898s   100000 91195.370/s
[GET] 1.173532594s   100000 85212.802/s
======================================
Msg:   16 N:   10000 Threads:   10 
[SET] 1.102275027s   100000 90721.460/s
[GET] 1.199071076s   100000 83397.892/s
```
