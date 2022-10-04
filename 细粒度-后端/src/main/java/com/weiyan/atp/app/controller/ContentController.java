package com.weiyan.atp.app.controller;

import com.weiyan.atp.data.bean.Result;
import com.weiyan.atp.data.request.web.DecryptContentRequest;
import com.weiyan.atp.data.request.web.ShareContentRequest;
import com.weiyan.atp.data.response.intergration.DecryptionResponse;
import com.weiyan.atp.data.response.intergration.EncryptionResponse;
import com.weiyan.atp.data.response.web.PlatContentsResponse;
import com.weiyan.atp.service.ContentService;

import lombok.extern.slf4j.Slf4j;

import org.apache.http.HttpEntity;
import org.apache.http.HttpResponse;
import org.apache.http.client.ClientProtocolException;
import org.apache.http.client.HttpClient;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.entity.StringEntity;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.util.EntityUtils;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.configurationprocessor.json.JSONArray;
import org.springframework.boot.configurationprocessor.json.JSONException;
import org.springframework.boot.configurationprocessor.json.JSONObject;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.io.IOException;
import java.nio.charset.StandardCharsets;

/**
 * @author : 魏延thor
 * @since : 2020/6/10
 */
@RestController
@RequestMapping("/content")
@Slf4j
public class ContentController {
    private final ContentService contentService;

//    @Value("${atp.devMode.baseUrl}")
//    private String baseUrl;

    public ContentController(ContentService contentService) {
        this.contentService = contentService;
    }

    @PostMapping("/")
    public Result<Object> encryptContent(@RequestBody @Validated ShareContentRequest request) {
        int tagSize = request.getTags().size();
        if (tagSize == 0 || tagSize > 10) {
            return Result.internalError("tags length error");
        }
        EncryptionResponse encryptionResponse = contentService.encContent(request);

//        String url = baseUrl+"/attrpolicy";
//        HttpClient client = HttpClients.createDefault();
//        //默认post请求
//        HttpPost post = new HttpPost(url);
//        //拼接多参数
//        JSONObject json = new JSONObject();
//        JSONArray array = new JSONArray();
//        try {
//            json.put("contentHash", encryptionResponse.getContentHash());
//            json.put("policy",request.getPolicy());
//            json.put("uid",request.getFileName());
//            array.put(request.getTags().get(0));
//            array.put(request.getTags().get(1));
//            array.put(request.getTags().get(2));
//            array.put(request.getTags().get(3));
//            json.put("tags",array);
//            json.put("timestamp",encryptionResponse.getTimeStamp());
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

        return Result.success();
    }

//    @PostMapping("/")
    @PostMapping("/share")
    public Result<Object> shareContent(@RequestBody @Validated ShareContentRequest request) {
        int tagSize = request.getTags().size();
        if (tagSize == 0 || tagSize > 10) {
            return Result.internalError("tags length error");
        }
        contentService.shareContent(request);
        return Result.success();
    }

    @PostMapping("/decryption")
    public Result<DecryptionResponse> decryptContent(@RequestBody @Validated DecryptContentRequest request) {
        String content = contentService.decryptContent(request.getCipher(), request.getFileName());
        DecryptionResponse response = DecryptionResponse.builder()
                .contentHash(request.getCipher())
                .uid(request.getFileName())
                .plainText(content)
                .tags(request.getTags())
                .timeStamp(String.valueOf(System.currentTimeMillis()))
                .build();

//        String url = baseUrl+"/share_judgement";
////        HttpClient client = HttpClients.createDefault();
////        //默认post请求
////        HttpPost post = new HttpPost(url);
////        //拼接多参数
////        JSONObject json = new JSONObject();
////        JSONArray array = new JSONArray();
////        try {
////            json.put("contentHash", request.getCipher());
////            json.put("plainText",content);
////            json.put("uid",request.getFileName());
////            array.put(request.getTags().get(0));
////            array.put(request.getTags().get(1));
////            array.put(request.getTags().get(2));
////            array.put(request.getTags().get(3));
////            json.put("tags",array);
////            json.put("timestamp",String.valueOf(System.currentTimeMillis()));
////            String message = "[" + json + "]";
////            post.addHeader("Content-type", "application/json; charset=utf-8");
////            post.setHeader("Accept", "application/json");
////            post.setEntity(new StringEntity(message, StandardCharsets.UTF_8));
////            HttpResponse httpResponse = client.execute(post);
////            HttpEntity entity = httpResponse.getEntity();
////            System.err.println("状态:" + httpResponse.getStatusLine());
////            System.err.println("参数:" + EntityUtils.toString(entity));
////
////        } catch (JSONException e) {
////            e.printStackTrace();
////        } catch (ClientProtocolException e) {
////            e.printStackTrace();
////        } catch (IOException e) {
////            e.printStackTrace();
////        }

        return Result.okWithData(response);
    }

    //@PostMapping("/decryption")
    @PostMapping("/decrypt")
    public Result<String> decContent(@RequestBody @Validated DecryptContentRequest request) {
        return Result.okWithData(
            contentService.decryptContent(request.getCipher(), request.getFileName()));
    }

    @GetMapping("/list")
    public Result<PlatContentsResponse> queryContents(String fromUserName, String tag,
                                                      int pageSize, String bookmark) {
        return Result.okWithData(
            contentService.queryPlatContents(fromUserName, tag, pageSize, bookmark));
    }
}
