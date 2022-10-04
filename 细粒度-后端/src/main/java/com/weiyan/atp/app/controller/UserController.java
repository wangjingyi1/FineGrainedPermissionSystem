package com.weiyan.atp.app.controller;

import com.weiyan.atp.constant.AttrApplyStatusEnum;
import com.weiyan.atp.constant.BaseException;
import com.weiyan.atp.data.bean.*;
import com.weiyan.atp.data.request.web.ApplyUserAttrRequest;
import com.weiyan.atp.data.request.web.ApproveAttrApplyRequest;
import com.weiyan.atp.data.request.web.CreateUserRequest;
import com.weiyan.atp.data.request.web.DeclareUserAttrRequest;
import com.weiyan.atp.data.request.web.QueryUserRequest;
import com.weiyan.atp.data.request.web.SyncAttrRequest;
import com.weiyan.atp.data.response.intergration.ApplyAttrResponse;
import com.weiyan.atp.data.response.intergration.CreateUserResponse;
import com.weiyan.atp.service.AttrService;
import com.weiyan.atp.service.UserRepositoryService;
import com.weiyan.atp.utils.JsonProviderHolder;

import lombok.extern.slf4j.Slf4j;

import org.apache.commons.io.FileUtils;
import org.apache.http.HttpEntity;
import org.apache.http.HttpResponse;
import org.apache.http.client.ClientProtocolException;
import org.apache.http.client.HttpClient;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.client.methods.HttpPut;
import org.apache.http.entity.StringEntity;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.util.EntityUtils;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.configurationprocessor.json.JSONException;
import org.springframework.boot.configurationprocessor.json.JSONObject;
import org.springframework.util.StringUtils;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.io.File;
import java.io.IOException;
import java.nio.charset.StandardCharsets;

/**
 * @author : 魏延thor
 * @since : 2020/6/1
 */
@RequestMapping("/user")
@RestController
@Slf4j
public class UserController {
    private final UserRepositoryService userRepositoryService;
    private final AttrService attrService;

    @Value("${atp.devMode.channelName}")
    private String channelName;

    @Value("${atp.path.cert}")
    private String certPath;

//    @Value("${atp.devMode.baseUrl}")
//    private String baseUrl;

    public UserController(UserRepositoryService userRepositoryService, AttrService attrService) {
        this.userRepositoryService = userRepositoryService;
        this.attrService = attrService;
    }

    @PostMapping("/")
    public Result<Object> createUser(@RequestBody @Validated CreateUserRequest request) {
        try {
            ChaincodeResponse response = userRepositoryService.createUser(request);
            String cert = FileUtils.readFileToString(new File(certPath + request.getFileName()),
                    StandardCharsets.UTF_8);
//            CreateUserResponse result = CreateUserResponse.builder()
//                    .cert(cert)
//                    .uid(request.getFileName())
//                    .channel(channelName)
//                    .timeStamp(String.valueOf(System.currentTimeMillis()))
//                    .result(response.isFailed() ? "failed" : "success")
//                    .build();
//            String url = baseUrl+"/attruser";
//            HttpClient client = HttpClients.createDefault();
//            //默认post请求
//            HttpPost post = new HttpPost(url);
//            //拼接多参数
//            JSONObject json = new JSONObject();
//            try {
//                json.put("channel_name", channelName);
//                json.put("certificate",cert);
//                json.put("uid",request.getFileName());
//                json.put("timestamp",String.valueOf(System.currentTimeMillis()));
//                String message = "[" + json + "]";
//                post.addHeader("Content-type", "application/json; charset=utf-8");
//                post.setHeader("Accept", "application/json");
//                post.setEntity(new StringEntity(message, StandardCharsets.UTF_8));
//                HttpResponse httpResponse = client.execute(post);
//                HttpEntity entity = httpResponse.getEntity();
//                System.err.println("状态:" + httpResponse.getStatusLine());
//                System.err.println("参数:" + EntityUtils.toString(entity));
//
//            } catch (JSONException e) {
//                e.printStackTrace();
//            }

            return Result.okWithData(response.getResult(str->str));
        } catch (IOException e) {
            e.printStackTrace();
            return null;
        }
    }

    @GetMapping("/")
    public Result<PlatUser> getUser(QueryUserRequest request) {
        if (StringUtils.isEmpty(request.getUserName()) && StringUtils.isEmpty(request.getPubKey())) {
            return Result.internalError("all empty request");
        }
        return Result.okWithData(userRepositoryService.queryUser(request));
    }

    @PostMapping("/attr")
    public Result<Object> declareAttr(@RequestBody @Validated DeclareUserAttrRequest request) {
        return attrService.declareUserAttr(request)
            .getResult(str -> str);
    }

    @PostMapping("/batchAttr")
    public Result<Object> batchDeclareAttr(@RequestBody @Validated DeclareUserAttrRequest request) {
        return attrService.batchDeclareUserAttr(request)
                .getResult(str -> str);
    }

    /**
     * 申请他人属性
     */
    @PostMapping("/attr/apply")
    public Result<Object> applyAttr(@RequestBody @Validated ApplyUserAttrRequest request) {
        ChaincodeResponse chaincodeResponse = attrService.applyAttr(request);
//        ApplyAttrResponse response = ApplyAttrResponse.builder()
//                .remark(request.getRemark())
//                .attrName(request.getAttrName())
//                .fromUserName(request.getToUserName())
//                .fromOrgName(request.getToOrgName())
//                .toUserName(request.getFileName())
//                .timeStamp(String.valueOf(System.currentTimeMillis()))
//                .result(chaincodeResponse.isFailed()?"failed":"success")
//                .build();

//        String url = baseUrl+"/addattr";
//        HttpClient client = HttpClients.createDefault();
//        //默认post请求
//        HttpPost post = new HttpPost(url);
//        //拼接多参数
//        JSONObject json = new JSONObject();
//        try {
//            json.put("channel_name", channelName);
//            json.put("fromUserName",request.getToUserName());
//            json.put("fromOrgName",request.getToOrgName());
//            json.put("toUserName",request.getFileName());
//            json.put("attrName",request.getAttrName());
//            json.put("timestamp",String.valueOf(System.currentTimeMillis()));
//            String message = "[" + json + "]";
//            post.addHeader("Content-type", "application/json; charset=utf-8");
//            post.setHeader("Accept", "application/json");
//            post.setEntity(new StringEntity(message, StandardCharsets.UTF_8));
//            HttpResponse httpResponse = client.execute(post);
//            HttpEntity entity = httpResponse.getEntity();
//            System.err.println("状态:" + httpResponse.getStatusLine());
//            System.err.println("参数:" + EntityUtils.toString(entity));
//
//        } catch (JSONException e) {
//            e.printStackTrace();
//        } catch (ClientProtocolException e) {
//            e.printStackTrace();
//        } catch (IOException e) {
//            e.printStackTrace();
//        }


        return Result.okWithData(chaincodeResponse.getResult(str->str));
    }

    /**
     * 申请他人属性
     */
//    @PostMapping("/attr/apply")
//    public Result<Object> applyAttr(@RequestBody @Validated ApplyUserAttrRequest request) {
//        return attrService.applyAttr(request)
//                .getResult(str -> str);
//    }

    /**
     * 批量申请他人属性
     */
    @PostMapping("/attr/batchApply")
    public Result<Object> batchApplyAttr(@RequestBody @Validated ApplyUserAttrRequest request) {
        return attrService.batchApplyAttr(request)
                .getResult(str -> str);
    }

    /**
     * 查询属性申请
     *
     * @param toId     查询的对方id
     * @param type     类型 0 for user； 1 for org
     * @param userName 用户名
     * @param status   状态
     */
    @GetMapping("/attr/apply")
    public Result<Object> queryAttrApply(String toId, Integer type, String userName,
                                         String status) {
        if (type != 0 && type != 1) {
            throw new BaseException("wrong type");
        }
        return attrService.queryAttrApply(type == 0 ? toId : "",
            type == 1 ? toId : "", userName, AttrApplyStatusEnum.valueOf(status))
            .getResult(str -> JsonProviderHolder.JACKSON.parseList(str, PlatUserAttrApply.class));
    }

    /**
     * 审批他人申请
     */
    @PostMapping("/attr/approval")
    public Result<Object> approveAttrApply(@RequestBody @Validated ApproveAttrApplyRequest request) {
        return attrService.approveAttrApply(request)
            .getResult(str -> str);
    }

    /**
     * 同步属性
     */
    @PostMapping("/attr/sync")
    public Result<DABEUser> syncSuccessApply(@RequestBody @Validated SyncAttrRequest request) {
        if (request.getType() == null) {
            return Result.okWithData(attrService.syncSuccessAttrApply(request.getFileName()));
        }
        return Result.okWithData(
            attrService.syncSuccessAttrApply(request.getFileName(),
                request.getType() == 0 ? request.getToId() : "",
                request.getType() == 1 ? request.getToId() : ""));
    }
}
