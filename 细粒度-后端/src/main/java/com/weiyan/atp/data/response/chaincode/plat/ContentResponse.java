package com.weiyan.atp.data.response.chaincode.plat;

import lombok.Data;

import java.util.List;

/**
 * @author : 魏延thor
 * @since : 2020/6/14
 */
@Data
public class ContentResponse {
    private String content;
    private String uid;
    private List<String> tags;
}
