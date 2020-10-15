# 读取以太坊稳定节点（DB里）信息

## 用法（widows）

1. 首先clone 代码到本地

2. 其次输入命令

   ```go
   main.exe "nodes路径"
   ```

   example：

   ```go
   main.exe  G:\work\区块链\以太坊\temp\nodes
   ```

3. 输出为：

   ```go
   Index: 97 IP: [50 116 55 109] ID: e426c80edc5a04a64c408b52d5117d34704bb2701a6a93aba5788b00d651186c Seq: 54 UDP: 30303
   Index: 98 IP: [104 248 55 223] ID: d4d5eb98ab5e10f9b3dc85ac59b6c4f28d64ccbb4a17455dc11157cfaae88eee Seq: 38 UDP: 3500
   Index: 99 IP: [195 68 24 189] ID: 8ad8c749800415589f922b1eea6fc1cac2445d89d225bfc00287e61711edea4a Seq: 0 UDP: 30303
   Index: 100 IP: [135 181 103 37] ID: e5c4e687ab1b451e226a61b4a17638dd212583e9fa63cb8126b4fba2cc3eb0f2 Seq: 1 UDP: 23000
   Index: 101 IP: [51 79 177 151] ID: 055047b2843d83db7b0752656ccf0c1af556940feecc66e8ce7173a74de4d3c0 Seq: 9 UDP: 30303
   Index: 102 IP: [94 130 22 36] ID: e4e2f990b819beb8c1e9d0cb6a725335c8032ac8b6c43a530bb6c10d74b4c755 Seq: 18 UDP: 30303
   ```

## 注意：Linux 请自己编译运行