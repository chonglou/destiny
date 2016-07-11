# destiny

命理计算库(阴阳、五行、八卦、农历等)

## 农历
中国的官方纪时采用的是中国公历（格里历），因此农历年历的推导应以公历年的周期为主导，附上农历年的信息，也就是说，年历以公历的1月1日为起始，至12月31日结束，根据农历历法推导出的农历日期信息，附加在公历日期信息上形成双历。通常情况下，一个公历年周期都不能完整地对应到一个农历年周期上，二者的偏差也不固定，因此不存在稳定的对应关系，也就是说，不存在从公历的日期到农历日期的转换公式，只能根据农历的历法规则推导出农历日期与公历日期的对应关系。
### 计算规则

### 节气

我国的农历的24节气是以视太阳位置定义的定气，不是以时间周期定义的平气。

它是这样规定的：

黄赤交点为春分点，并定为黄经0度，黄道一周为360度。

- 黄经度数能被30整除的为中气，共12个；
- 黄经度数能被15整除但不能被30整除的为节气，也有12个；
- 节气和中气合称为节气，共24个；
- 黄经度数能被5整除的为定候，共有72个，其中有24个为节气，其余的为一般的定候；

所以在日历上就得标注节气的发生时间，这由天文台来计算的。

### 虚岁与周岁

虚岁是中华民族一种重要的民俗，它是以春节（农历新年）分界点的，而周岁以本人的生日为分界点的；

虚岁换算成周岁
- 在本人生日到来之前，周岁=虚岁-2

- 在本人生日到来及以后，周岁=虚岁-1；

周岁换算成虚岁
- 在本人生日到来之前，虚岁=周岁+2；
- 在本人生日到来及以后，虚岁=周岁+1；

## 参考文档

- <http://sinocal.sinica.edu.tw/>
- <http://blog.csdn.net/orbit/article/details/9337377>
- <http://ccal.chinesebay.com/ccal/ccal.htm.cn>
