# 一、系统接口
## 1.注册接口：

- 请求地址：
url: http://www.doit05.top:9020/system/register

- 请求方法：post

- 必要参数：
    - name:admin
    - password:admin123
    - email:768068275@qq.com
    - mobile:18520142853
- 请求返回值：

    -  请求成功：
    { "status": "0000", "msg": "success", "data": [ 5 ] }
    -  请求失败：
    { "status": -6, "msg": "参数缺失", "data": {} }

- 数据表：user_basic
- 备注：

其他参数请参照数据表


## 2.登录接口

- 请求地址：
url: http://www.doit05.top:9020/system/login

- 请求方法：post

- 必要参数：
    - name:admin
    - password:admin123
- 请求返回值：
    -  请求成功：
    { "status": "0000", "msg": "success", "data": [ 5 ] }
    -  请求失败：
    { "status": -6, "msg": "参数缺失", "data": {} }

- 数据表：user_basic
- 备注：

    其他参数请参照数据表


## 3.首页广告接口

- 请求地址：
url: http://www.doit05.top:9020/system/advertisments

- 请求方法：GET

- 必要参数：
    - name:admin
    - password:admin123
- 请求返回值：
    -  请求成功：
    { "status": "0000", "msg": "success", "data": [ url0,url1,url2,url3 ] }==(url0,url1,url2,url3为广告图片网址)==
    -  请求失败：
    { "status": -6, "msg": "fail", "data": {} }

- 数据表：==user_basic==
- 备注：

    其他参数请参照数据表

## 4.数据监控（柱图）接口
1>能量摄入/消耗（首页左边柱图）

请求地址：

url:http://119.29.109.120:9020/barchart/left

参数：
日期 String date="2016-06-03";
请求方式：POST

###返回值：

请求接口失败 { "status": -6, "msg": "参数缺失", "data": {} }

成功并有返回数据 { "status": "0000", "msg": "success", "data": [{“wastage”:2000.0}， {"intake":1600.0}, {"normal":1700.0}] }

返回数据说明

消耗能量 float wastage（基础消耗能量+运动消耗能量）
摄入能量 float intake(早中晚餐摄入总能量)
摄入能量标准值 float normal（根据年龄段区分）
返回数据单位为千卡路里（kcal）
2>监控指标数据（首页右边柱图）

请求地址：

url:http://119.29.109.120:9020/barchart/right

参数：
日期 String date="2016-06-03"; （监控指标有多个，此处以体重举例）
###返回值：

请求接口失败 { "status": -6, "msg": "参数缺失", "data": {} }

成功并有返回数据 { "status": "0000", "msg": "success", "data": [{“weight”:65.0}， {"normal":60.0}] }

返回数据说明

体重 float weight
摄入能量标准值 float normal
体重单位为KG
由于监控数据源不清楚，可以在需求讨论清楚再做
五、监控接口
六、食品录入接口
请求地址：

url:http://119.29.109.120:9020/barchart/inputFood

###返回值：

请求接口失败 { "status": -6, "msg": "fail", "data": {} }

成功并有返回数据 { "status": "0000", "msg": "success", "data": [ { “sumIntake”: “2500.0”, "breakfastIntake": “600.0”, "lunchIntake": “1000.0”, "dinnerIntake": “900.0” }, { "breakfastDiet": [ { "foodName"："鸡蛋"，"data": [ { "energy": 200.0, "url": "httP://", "weight": "200", "unit": "个" } ] }，{ "foodName"："面包"，"data": [ { "energy": 200.0, "url": "httP://", "weight": "200", "unit": "克" } ] } ], "lunchDiet": [ { "foodName"："鸡蛋"，"data": [ { "energy": 200.0, "url": "httP://", "weight": "200", "unit": "个" } ] } ], "dinnerDiet": [ { "foodName"："鸡蛋"，"data": [ { "energy": 200.0, "url": "httP://", "weight": "200", "unit": "个" } ] } ] }

查询食品接口
请求地址：

url:http://119.29.109.120:9020/findFood

请求方式 POST
请求参数
int flag(0、1、2，分别代表早、中、晚餐)
返回值

请求接口失败 { "status": -6, "msg": "fail", "data": {} } -成功并返回数据 { "status": "0000", "msg": "success", "data": [ { "foodClassify":"产妇菜谱"， "foodName": "鸡蛋" }, { "foodClassify":"产妇菜谱"， "foodName": "鸡蛋" }, { "foodClassify":"产妇菜谱"， "foodName": "鸡蛋" } ] }
返回数据说明：

根据早中晚餐返回相应食物，每次请求返回10-15条数据即可
食品详情接口
请求地址：

url:http://119.29.109.120:9020/findFoodDetail

请求方式 POST
请求参数 1.String foodName
返回值

请求接口失败 { "status": -6, "msg": "fail", "data": {} } -成功并返回数据 { "status": "0000", "msg": "fail", "isSuggest":0 "data":{ "foodName":"鸡蛋"， "energy":54.3, "fat":17.3, ... } }
返回数据说明

返回数据主要返回食品营养成分，具体参考数据库关于foodnutri的几张表
一般数据单位都默认（重量为克，能量为千卡等等），所以可以不用传到前端，若有特殊说明则可以用一键值对表示，例如： "energyUnit":"卡"
isSuggest(是否推荐)，0表示不推荐，1表示推荐。这个数据需要根据用户信息来计算。
添加食物接口
请求地址：

url:http://119.29.109.120:9020/addFood

请求方式 POST
请求参数
String foodName
String weight
String energy
int isCollect(0表示未收藏，1表示已收藏)
返回值

请求接口失败 { "status": -6, "msg": "fail" } -请求成功并返回数据 { "status": "0000", "msg": "success" }
账户信息接口
请求地址：

url:http://119.29.109.120:9020/addFood

请求方式 POST
请求参数
帐号（手机号或用户名），对应两个参数
String name
String mobile
返回值

请求接口失败 { "status": -6, "msg": "fail" "data":[] } -请求成功并返回数据 { "status": "0000", "msg": "success" "data":{ "userName":"bob", "gender":"男"， "email":"747184473@qq.com", ...... } }


# 二、基础数据接口
## 1.设置系统目标接口：
- 请求地址：
url: http://www.doit05.top:9020/system/setgoal

- 请求方法：post

- 必要参数：
    - name:减肥w
    - imgurl : /img/a.jpg

- 请求返回值：

    -  请求成功：
    { "status": "0000", "msg": "success", "data": [4] }
    -  请求失败：
    { "status": -6, "msg": "参数缺失", "data": {} }

- 数据表：user_goal
- 备注：

    imgurl为绝对路径
    name唯一

## 2.获取系统目标接口：

- 请求地址：
url: http://www.doit05.top:9020/system/getgoals

- 请求方法：get

- 必要参数：

- 请求返回值：

    -  请求成功：
    { "status": "0000", "msg": "success", "data": [] }
    -  请求失败：
    { "status": -6, "msg": "参数缺失", "data": {} }

- 数据表：user_goal
- 备注：

## 3.删除系统目标接口：
- 请求地址：
url: http://www.doit05.top:9020/system/deletegoal

- 请求方法：post

- 必要参数：
    - name:减肥w
    - imgurl : /img/a.jpg

- 请求返回值：

    -  请求成功：
    { "status": "0000", "msg": "success", "data": [4] }
    -  请求失败：
    { "status": -6, "msg": "参数缺失", "data": {} }

- 数据表：user_goal
- 备注：

    imgurl为绝对路径
    name唯一


## 4.用户添加目标接口

- 请求地址：
url: http://www.doit05.top:9020/basic/choosegoal

- 请求方法：post

- 必要参数：
    - uid:3
    - name:减肥
- 请求返回值：
    -  请求成功：
    { "status": "0000", "msg": "success", "data": [ 5 ] }
    -  请求失败：
    { "status": -6, "msg": "参数缺失", "data": {} }

- 数据表：user_goal
- 备注：

    其他参数请参照数据表

## 5.用户删除目标接口

    - 请求地址：
    url: http://www.doit05.top:9020/basic/deletegoal

    - 请求方法：post

    - 必要参数：
        - uid:3
        - name:减肥
    - 请求返回值：
        -  请求成功：
        { "status": "0000", "msg": "success", "data": [ 5 ] }
        -  请求失败：
        { "status": -6, "msg": "参数缺失", "data": {} }

    - 数据表：user_goal
    - 备注：

        其他参数请参照数据表

## 6.获取用户目标接口

- 请求地址：
url: http://www.doit05.top:9020/basic/getusergoals?==uid=?==

- 请求方法：get

- 必要参数：
    - uid:2
- 请求返回值：
    -  请求成功：
    {
        -   "status": "0000",
        -   "msg": "success",
        -   "data": [
    {
      "Id": 1,
      "Uid": 1,
      "Goal_name": "减肥",
      "Goal_switch": 1,
      "Start": 0,
      "End": 0,
      "Weekly_change": 0,
      "Increase": 0,
      "Goal_date": 0,
      "Daily_entry": 0,
      "Daily_foodpoint": 0,
      "Create_time": 1466220667,
      "Update_time": 1466220667
    }
  ]
}
    -  请求失败：
    { "status": -6, "msg": "参数缺失", "data": {} }

- 数据表：user_goal
- 备注：

    其他参数请参照数据表

## 7.设置系统监控接口：
    - 请求地址：
    url: http://www.doit05.top:9020/system/setmonitor

    - 请求方法：post

    - 必要参数：
        - name:减肥w
        - imgurl : /img/a.jpg

    - 请求返回值：

        -  请求成功：
        { "status": "0000", "msg": "success", "data": [4] }
        -  请求失败：
        { "status": -6, "msg": "参数缺失", "data": {} }

    - 数据表：user_goal
    - 备注：

        imgurl为绝对路径
        name唯一

## 8.获取系统监控接口：

- 请求地址：
url: http://www.doit05.top:9020/basic/getsystemmonitors

- 请求方法：get

- 必要参数：

- 请求返回值：

    -  请求成功：
    { "status": "0000", "msg": "success", "data": [] }
    -  请求失败：
    { "status": -6, "msg": "参数缺失", "data": {} }

- 数据表：user_monitor
- 备注：

    uid = 1
## 9.删除系统监控接口：
    - 请求地址：
    url: http://www.doit05.top:9020/system/deletemonitor

    - 请求方法：post

    - 必要参数：
        - name:减肥w

    - 请求返回值：

        -  请求成功：
        { "status": "0000", "msg": "success", "data": [4] }
        -  请求失败：
        { "status": -6, "msg": "参数缺失", "data": {} }

    - 数据表：user_goal
    - 备注：
        name唯一

## 10.用户选择监控接口
        - 请求地址：
        url: http://www.doit05.top:9020/basic/choosemonitor

        - 请求方法：post

        - 必要参数：
            - uid:3
            - name:减肥
        - 请求返回值：
            -  请求成功：
            { "status": "0000", "msg": "success", "data": [ 5 ] }
            -  请求失败：
            { "status": -6, "msg": "参数缺失", "data": {} }

        - 数据表：user_goal
        - 备注：

## 6.获取用户监控接口

        - 请求地址：
        url: http://www.doit05.top:9020/basic/getusermonitors?==uid=?==

        - 请求方法：get

        - 必要参数：
            - uid:2
        - 请求返回值：
            -  请求成功：
            {
                -   "status": "0000",
                -   "msg": "success",
                -   "data": [
            {
              "Id": 1,
              "Uid": 1,
              "name": "减肥",
              "switch": 1,
              "Start": 0,
              "End": 0,
              "Weekly_change": 0,
              "Increase": 0,
              "Date": 0,
              "Daily_entry": 0,
              "Daily_foodpoint": 0,
              "Create_time": 1466220667,
              "Update_time": 1466220667
            }
          ]
        }
            -  请求失败：
            { "status": -6, "msg": "参数缺失", "data": {} }

        - 数据表：user_goal
        - 备注：

            其他参数请参照数据表
## 11.用户监控数据接口

- 请求地址：
url: http://www.doit05.top:9020/basic/upmonitordata

- 请求方法：post

- 必要参数：
    - uid : 2
    - name : "脂肪"
    - value : 100
    - valuetime :1466220667
- 请求返回值：
    -  请求成功：
    { "status": "0000", "msg": "success", "data": [ 5 ] }
    -  请求失败：
    { "status": -6, "msg": "参数缺失", "data": {} }

- 数据表：user_monitor
- 备注：

    其他参数请参照数据表

## 6.获取用户监控接口

- 请求地址：
url: http://www.doit05.top:9020/basic/getusermonitors?==uid=?==

- 请求方法：get

- 必要参数：
    - uid:1
    - type : 2
- 请求返回值：
    -  请求成功：
    {
        -   "status": "0000",
        -   "msg": "success",
        -   "data": [
    {
      "Id": 1,
      "Uid": 1,
      "Monitor_name": "减肥",
      "Monitor_switch": 1,
      "Value": 0,
      "Value_time": 0,
      "Create_time": 1466220667,
      "Update_time": 1466220667
    }
  ]
}
    -  请求失败：
    { "status": -6, "msg": "参数缺失", "data": {} }

- 数据表：user_monitor
- 备注：

    其他参数请参照数据表

#     三、用户日常数据接口
# 四、系统知识库接口
