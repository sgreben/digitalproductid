# digitalproductid

Decodes the Windows 7/8/10 license key from its binary registry encoding.

## Get it

```text
go get -u github.com/sgreben/digitalproductid
```

## Usage

```sh
# Win7/8
c:\> digitalproductid
XXXXX-XXXXX-XXXXX-XXXXX-XXXXX
```

```sh
# Win10
c:\> digitalproductid -n DigitalProductId
XXXXX-XXXXX-XXXXX-XXXXX-XXXXX
```
