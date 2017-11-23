# az-dc-ranges
Azure DC IP Ranges parser.

This tool downloads the latest XML file from [https://www.microsoft.com/en-us/download/confirmation.aspx?id=41653](https://www.microsoft.com/en-us/download/confirmation.aspx?id=41653) and returns valid JSON output.

Check the source code to easily change the above for China's region ranges.

## Disclaimer
THIS IS PROVIDED AS-IS. IT MIGHT STOP WORKING AT ANY POINT IN TIME WITHOUT PREVIOUS NOTICE.    
THIS IS NOT OFFICIAL MICROSOFT SOFTWARE. THIS IS THE WORK OF AN INDIVIDUAL ON THEIR SPARE TIME.

### Usage
```
~$ git clone https://github.com/PedroPerezMSFT/az-dc-ranges
~$ cd az-dc-ranges
~$ go build main.go
~$ ./main
```

### Output (excerpt):

```
{"Region":[{"Name":"europewest","IPRange":[{"Subnet":"40.112.124.0/24"},{"Subnet":"65.52.128.0/19"},{"Subnet":"94.245.97.0/24"},{"Subnet":"104.47.169.0/24"},{"Subnet":"104.214.240.0/24"},{"Subnet":"137.116.192.0/19"},{"Subnet":"168.63.0.
0/19"},{"Subnet":"168.63.96.0/20"},{"Subnet":"168.63.112.16/28"},{"Subnet":"168.63.112.64/26"},{"Subnet":"168.63.112.128/25"},{"Subnet":"168.63.113.0/24"},{"Subnet":"168.63.114.0/23"},{"Subnet":"168.63.116.0/22"},{"Subnet":"168.63.120.0/
21"},{"Subnet":"193.149.80.0/22"},{"Subnet":"213.199.128.0/21"},{"Subnet":"213.199.136.0/22"},{"Subnet":"213.199.180.32/28"},{"Subnet":"213.199.180.112/28"},{"Subnet":"213.199.183.0/24"},{"Subnet":"23.97.128.0/17"},{"Subnet":"23.98.46.0/
24"},{"Subnet":"23.100.0.0/20"},{"Subnet":"23.101.64.0/20"},{"Subnet":"40.74.0.0/18"},{"Subnet":"40.90.141.160/27"},{"Subnet":"40.114.128.0/18"},
```

### Working with the output
#### Linux (jq)
```
~$ ./main | jq 
{
  "Region": [
    {
      "Name": "europewest",
      "IPRange": [
        {
          "Subnet": "40.112.124.0/24"
        },
        {
          "Subnet": "65.52.128.0/19"
        },
        {
          "Subnet": "94.245.97.0/24"
        },
        {
          "Subnet": "104.47.169.0/24"
        },
        {
          "Subnet": "104.214.240.0/24"
        },
        {
          "Subnet": "137.116.192.0/19"
        },
        {
          "Subnet": "168.63.0.0/19"
        },
        {
          "Subnet": "168.63.96.0/20"
        },
        {
          "Subnet": "168.63.112.16/28"
        },
        {
          "Subnet": "168.63.112.64/26"
        },
        {
          "Subnet": "168.63.112.128/25"
        },
        {
          "Subnet": "168.63.113.0/24"
        },
        {
          "Subnet": "168.63.114.0/23"
        },
        {
          "Subnet": "168.63.116.0/22"
        },
```

#### Windows (PowerShell) (**Also works on Linux with [PowerShell Core](https://github.com/PowerShell/PowerShell)!!)
```
~$ $results = ./main.exe | ConvertFrom-JSON
~$ $results.Region[0].Name; $results.Region[0].IPRange
europewest

Subnet
------
40.112.124.0/24
65.52.128.0/19
94.245.97.0/24
104.47.169.0/24
104.214.240.0/24
137.116.192.0/19
168.63.0.0/19
168.63.96.0/20
168.63.112.16/28
168.63.112.64/26
168.63.112.128/25
168.63.113.0/24
168.63.114.0/23
168.63.116.0/22
(...)
```

