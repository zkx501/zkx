/**
 * Created by Administrator on 2016/5/20.
 */
// 对Date的扩展，将 Date 转化为指定格式的String
    // 月(M)、日(d)、小时(h)、分(m)、秒(s)、季度(q) 可以用 1-2 个占位符，
    // 年(y)可以用 1-4 个占位符，毫秒(S)只能用 1 个占位符(是 1-3 位的数字)
    // 例子：
    // (new Date()).Format("yyyy-MM-dd hh:mm:ss.S") ==> 2006-07-02 08:09:04.423
    // (new Date()).Format("yyyy-M-d h:m:s.S")      ==> 2006-7-2 8:9:4.18
Date.prototype.dateFormat = function(fmt,year)
{ //author: meizz
    var o = {
        "M+" : this.getMonth()+1,                 //月份
        "d+" : this.getDate(),                    //日
        "h+" : this.getHours(),                   //小时
        "m+" : this.getMinutes(),                 //分
        "s+" : this.getSeconds(),                 //秒
        "q+" : Math.floor((this.getMonth()+3)/3), //季度
        "S"  : this.getMilliseconds()             //毫秒
    };
    var now_year = this.getFullYear();
    if(!isNaN(year)){
        now_year += year;
    }
    if(/(y+)/.test(fmt))
        fmt=fmt.replace(RegExp.$1, (now_year+"").substr(4 - RegExp.$1.length));
    for(var k in o)
        if(new RegExp("("+ k +")").test(fmt))
            fmt = fmt.replace(RegExp.$1, (RegExp.$1.length==1) ? (o[k]) : (("00"+ o[k]).substr((""+ o[k]).length)));
    return fmt;
};

var debug = true;
function log(str){
    if(debug){
        console.log(str);
    }
}

function logo(obj,prefix,suffix){
    if(debug){
        console.log((prefix?prefix:'')+JSON.stringify(obj)+(suffix?suffix:''));
    }
}

var info = true;
function logInfo(str){
    if(info){
        console.info(str);
    }
}

function logoInfo(obj,prefix,suffix){
    if(info){
        console.info((prefix?prefix:'')+JSON.stringify(obj)+(suffix?suffix:''));
    }
}

/**
 * 通用的数据交互方法
 * @param url	请求url地址
 * @param param		请求参数
 * @param callback	回调方法
 * @param success_msg	成功操作时要显示的信息
 * @param null_msg	无结果返回要显示的信息
 */
function webData(url,param,callback,success_msg,null_msg){
    $.ajax({
        type: "POST",
        url: url,
        cache : false,	//禁用缓存
        data: param,	//传入已封装的参数
        dataType: "json",
        success: function(result) {
            console.log("ajax结果："+JSON.stringify(result));
            if(typeof result==undefined || result==null){
                if(null_msg==undefined || null_msg==null || null_msg.length==0){
                    null_msg = "无结果返回";
                }
                layer.alert(null_msg, {icon: 2,title:'<span style="color: red">错误</span>'});
                return;
            }
            CodeValidator.validator(result,function(data){
                layer.alert("错误码"+data.code,{icon: 2,title:'<span style="color: red">错误</span>'});
                console.info("错误信息"+data.message);
            },function(data){
                if(typeof success_msg!=undefined && success_msg!=null && success_msg.length>0){
                    layer.msg(success_msg,{
                        time:1500,
                    });
                }
                if(typeof callback == 'function'){
                    callback(result);
                }
            });
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            layer.alert('系统错误', {icon: 2,title:'<span style="color: red">错误</span>'});
        }
    });
}




//获取省市县乡村等信息
function get_area_info(obj,area_condition){
    $.ajax({
        url: Url.manage+'area/showCommonAreaAll',
        type:'post',
        data:area_condition,
        dataType:'jsonp',
        success:function(result){
            //其实这里要对result是否为空判断
            var rows = result.data;
            if(rows && rows.length>0){
                for(var i in rows){
                    $('#'+obj).append('<option value="'+rows[i].id+'">'+rows[i].name+'</option>');
                }
            }
        }
    });
}

//省变化时，市，区，镇都清空，并重新加载市信息
$('#province').on('change',function(){
    $('#gonghuoAddr').val('');
    $('#city,#district,#township').empty().html('<option value="">--请选择--</option>');
    get_area_info('city',{pid:$('#province').val(),level:2});
});

//市变化时，区，镇都清空，并重新加载区信息
$('#city').on('change',function(){
    $('#gonghuoAddr').val('');
    $('#district,#township').empty().html('<option value="">--请选择--</option>');
    get_area_info('district',{pid:$('#city').val(),level:3});
});

//区变化时，镇清空，并重新加载镇信息
$('#district').on('change',function(){
    $('#gonghuoAddr').val('');
    $('#township').empty().html('<option value="">--请选择--</option>');
    get_area_info('township',{pid:$('#district').val(),level:4});
});

//镇变化时，街道清空，并重新加载街道信息
$('#township').on('change',function(){
    $('#gonghuoAddr').val('');
    $('#committee').empty().html('<option value="">--请选择--</option>');
    get_area_info('committee',{pid:$('#township').val(),level:5});
});

var getAddressDetail = function(){
    return {
        provinceName : $('#province option:selected').text(),
        provinceCode : $('#province option:selected').val(),
        cityName : $('#city option:selected').text(),
        cityCode : $('#city option:selected').val(),
        districtName : $('#district option:selected').text(),
        districtCode : $('#district option:selected').val(),
        townshipName : $('#township option:selected').text(),
        townshipCode : $('#township option:selected').val(),
        committeeName : $('#committee option:selected').text(),
        committeeCode : $('#committee option:selected').val(),
    };
};

//本方法未测试过
//var removeChar = function(char){
//    return char.replace(/-/g,"");
//};

//日历组件
//        $('.single_cal').daterangepicker({
//            singleDatePicker: true,
//            calender_style:"picker_4",
//            //minDate: '2012/01/01',
//            format:'YYYY/MM/DD',
//            locale:{
//                applyLabel:'Submit',
//                daysOfWeek: ['日', '一', '二', '三', '四', '五','六'],
//                monthNames: ['一月', '二月', '三月', '四月', '五月', '六月', '七月', '八月',
//                    '九月', '十月', '十一月', '十二月']
//            }
//        },function(start, end, label){
//            console.info("");
//        });

//validate_area为需要验证区域的 id/class/html标签,不填则验证整个html的input/select
//init是对required的input和select进行失去焦点判断
//isTrue在form表单提交时，验证有required的input和select是否为空
var jsb_validate = function(){
    return {
        init:function(validate_area){
            if(!validate_area){
                validate_area = '';
            }
            $(validate_area+' input').iCheck({
                checkboxClass: 'icheckbox_flat-green',
                radioClass: 'iradio_flat-green'
            });
            $(validate_area+' input[required]').on('blur',function(){
                if($(this).val()==""){
                    new PNotify({
                        'title':'警告:',
                        'text':$(this).data('error-msg')+'不能为空! ',
                        'type':'warning',
                        'delay':2000
                    });
                }
            });

            $(validate_area+' select[required]').on('change',function(){
                if($(this).find('option:selected').val()==""){
                    new PNotify({
                        'title':'警告:',
                        'text':$(this).data('error-msg')+'不能为空! ',
                        'type':'warning',
                        'delay':2000
                    });
                }
            });
        },
        isTrue:function(validate_area){
            if(!validate_area){
                validate_area = '';
            }
            var flag = true;
            $(validate_area+' input[required]').each(function(){
                console.log($(this).data('error-msg')+" "+$(this).val());
                if($(this).val()==""){
                    new PNotify({
                        'title':'警告:',
                        'text':$(this).data('error-msg')+'不能为空! ',
                        'type':'warning',
                        'delay':3000
                    });
                    flag = false;
                    //return false;
                }
            });
            $(validate_area+' select[required]').each(function(){
                console.log($(this).data('error-msg')+" "+$(this).find('option:selected').val());
                if($(this).find('option:selected').val()==""){
                    new PNotify({
                        'title':'警告:',
                        'text':$(this).data('error-msg')+'不能为空! ',
                        'type':'warning',
                        'delay':3000
                    });
                    flag = false;
                    //return false;
                }
            });
            return flag;
        }
    }
}();

/**
 * 字符串格式化
 * @param str       需要格式化的字符串   ex:123456789
 * @param times     间隔  ex:4
 * @param separator 分隔符 ex:" "
 * @returns {string}    1234 5678 9
 */
function str_format(str,times,separator){
    str+='';
    var result = '';
    while(str.length>times){
        result = result+str.substring(0,times)+separator;
        str = str.substring(times);
    }
    result += str;
    return result;
}
