package com.weiyan.atp.data.bean;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.weiyan.atp.data.response.chaincode.plat.ContentResponse;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * @author : 魏延thor
 * @since : 2020/6/11
 */
@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class PlatContent {
    private String policy;

    private String cipher;

    private List<String> tags;

    @JsonIgnore
    private static final String PATTERN = "Policy\":\"(.*?)\"";

    public PlatContent(ContentResponse response) {
        this.tags = response.getTags();
        this.cipher = response.getContent();
        Matcher matcher = Pattern.compile(PATTERN).matcher(cipher);
        if (matcher.find() && matcher.groupCount() == 1) {
            this.policy = matcher.group(1);
        }
    }
}
