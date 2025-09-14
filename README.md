<h1 align="center">TinyFetch</h1>

<p align="center">
<img src="https://img.shields.io/github/license/steeldregg/tinyfetch?color=blue&style=plastic">
<img src="https://img.shields.io/badge/language-Go-blue?style=plastic">
<img src="https://img.shields.io/badge/stauts-stable-green?style=plastic">
</p>

TinyFetch is a system info tool written in GoLang

## Showcase

Ubuntu
```text
dregg@server:~$ ./tinyfetch-linux-amd64

              )sSQQQQQQQQQQso                  dregg@server
          )SQQQQQQQQQQQQQQQQQQQQo                OS: ubuntu 22.04 x86_64
        SQQQQQQQQQQQQQQQQQbCCPSQQQb          Kernel: linux 5.15.0-140-generic
      SQQQQQQQQQQQSbbbbbGb     QQQQQb           CPU: Intel(R) Xeon(R) Gold 6150 CPU @ 2.70GHz (1x36)
     QQQQQQQQQbGQ        Pp  )QQQQQQQQp         RAM: 907.7 MiB / 62.8 GiB (1%)
    QQQQQQQQb    Q)spQpso     PQQQQQQQQp       Swap: 0 B
   QQQQQQQQb    QQQQQQQQQQQo    SQQQQQQQ       Disk: 14.5 GiB / 30.3 GiB (50%)
  )QQQSbbbS   (QQQQQQQQQQQQQp   )QQQQQQQb  
  )QQS     Q  QQQQQQQQQQQQQQGssspQQQQQQQb  
  )QQQp   sb  GQQQQQQQQQQQQQS   )QQQQQQQb  
   QQQQQQQQ    GQQQQQQQQQQQb    QQQQQQQQb  
   )QQQQQQQQ     SQQQQQQSb     QQQQQQQQb   
    PQQQQQQQQb  Qb        sDGSQQQQQQQQb    
     )QQQQQQQQQQQpo     (b    )QQQQQQb     
       PQQQQQQQQQQQQQQQQQQ    sQQQQb       
         PSQQQQQQQQQQQQQQQQQQQQQSb         
            (PSQQQQQQQQQQQQQSbC            
                   (CDCC                   
                   
```

MacOs
```text
dregg@DrEggs-Mac-Pro ~ % ./tinyfetch-darwin-amd64

                         )s                 Host: DrEggs-Mac-Pro.local
                      sQQQQ                   OS: darwin 12.7.1 x86_64
                    )QQQQQb               Kernel: darwin 21.6.0
                    QQQQb                    CPU: Intel(R) Xeon(R) W-3235 CPU @ 3.30GHz (12)
            )oo     PP)  oppo                RAM: 16.8 GiB / 192.0 GiB (9%)
        sQQQQQQQQQQQSQQQQQQQQQQQp           Swap: 0 B
      )QQQQQQQQQQQQQQQQQQQQQQQQQbD           GPU: AMD Radeon Pro W5700X
     (QQQQQQQQQQQQQQQQQQQQQQQQSb            Disk: 656.7 GiB / 953.7 GiB (69%)
     QQQQQQQQQQQQQQQQQQQQQQQQQb         
    )QQQQQQQQQQQQQQQQQQQQQQQQQ          
     QQQQQQQQQQQQQQQQQQQQQQQQQo         
     QQQQQQQQQQQQQQQQQQQQQQQQQQp        
     )QQQQQQQQQQQQQQQQQQQQQQQQQQSp      
      )QQQQQQQQQQQQQQQQQQQQQQQQQQQb     
       )QQQQQQQQQQQQQQQQQQQQQQQQQb      
        (QQQQQQQQQQQQQQQQQQQQQQSb       
          )SQQQQQQSSbSSQQQQQQSb         
             PDb          PP)           
             
```

Windows
```text
PS C:\Users\Administrator> .\tinyfetch-windows-amd64.exe

         @@@@@@@@@@@@@@@                      Host: WIN-SERVER
        @@@@@@@@@@@@@@@   @@                    OS: Microsoft Windows Server 2022 Datacenter 21H2 x86_64
       @@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@   Kernel: windows 10.0.20348.2762 Build 20348.2762
       @@@@@@@@@@@@@@@   @@@@@@@@@@@@@@@       CPU: Intel(R) Xeon(R) Gold 6150 CPU @ 2.70GHz (36x2)
       @@@@@@@@@@@@@@    @@@@@@@@@@@@@@@       RAM: 102.4 GiB / 766.7 GiB (13%)
      @@@@@@@@@@@@@@@   @@@@@@@@@@@@@@@@      Swap: 0 B
     @            @@    @@@@@@@@@@@@@@@        GPU: NVIDIA Tesla P4
                       @@@@@@@@@@@@@@@        Disk: 49.7 GiB / 222.9 GiB (22%)
     @@@@@@@@@@@@@      @@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@         @@@@@
   @@@@@@@@@@@@@@@    @
   @@@@@@@@@@@@@@@   @@@@@@@@@@@@@@@
  @@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@
  @@@@@@@@@@@@@@@   @@@@@@@@@@@@@@@
 @@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@
             @@@   @@@@@@@@@@@@@@@
                   @@@@@@@@@@@@@@@

```