# server layer design

## parser

### RESP

> client与server间采用TCP协议进行传输，在传输层之上，使用RESP协议解决TCP粘包拆包的问题。

- 从network layer拿到的原始数据是 `[]byte` 类型，parser需要通过RESP协议解析成相应的RESP报文。

- parser将原始数据解析成RESP报文后，交由executor执行。

- 将executor执行的结果通过network layer回传给client

reply type：

<details><summary>details</summary>

- simple strings:

  > 非二进制安全，但可以最快地传输字符串，例如OK

  +OK\r\n

- integers:

  :1\r\n

- errors:

  -invalid command\r\n

- bulk string:

  > 二进制安全，通过附带消息头解决二进制安全问题

  $11\r\nhello world\r\n

  $-1\r\n  => null

- array:

  > 数组的元素可以是不同类型的

  *3\r\n$3\r\nset\r\n$3\r\nkey\r\n$5\r\nvalue\r\n

  *3\r\n:1\r\n:2\r\n$3\r\nval\r\n

</details>

## executor
