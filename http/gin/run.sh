#!/bin/bash

# 测试 AsciiJSON 特性
curl -v http://127.0.0.1:8080/ascii/json

# 测试 HTML render 特性
curl -v http://127.0.0.1:8080/posts/index
curl -v http://127.0.0.1:8080/users/index

# 测试 jsonp 特性
curl -v http://127.0.0.1:8080/jsonp?callback=x

# 测试 Mutlipart/Urlencoded 绑定特性
curl -v --form user=user --form password=password http://127.0.0.1:8080/login
curl -v --form message=这是一个表单消息 --form nick=未知 http://localhost:8080/form_post

# 测试 PureJSON 特性
curl -v http://127.0.0.1:8080/json
curl -v http://127.0.0.1:8080/purejson

# 测试 Query and PostForm 特性
curl -v --form name=manu --form message=this_is_great -H "Content-Type: application/x-www-form-urlencoded" \
    http://127.0.0.1:8080/post?id=1234&page=1

# 测试 SecureJSON 特性
curl -v http://127.0.0.1:8080/securejson