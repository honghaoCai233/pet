package clients

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"strings"

	"pet/configs"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OSSClient struct {
	bucket     *oss.Bucket
	endpoint   string
	bucketName string
}

// NewOSSClient 创建OSS客户端
func NewOSSClient(conf *configs.Config) (*OSSClient, error) {
	client, err := oss.New(conf.OSS.Endpoint, conf.OSS.AccessKeyId, conf.OSS.AccessKeySecret)
	if err != nil {
		return nil, fmt.Errorf("创建OSS客户端失败: %v", err)
	}

	bucket, err := client.Bucket(conf.OSS.BucketName)
	if err != nil {
		return nil, fmt.Errorf("获取Bucket失败: %v", err)
	}

	return &OSSClient{
		bucket:     bucket,
		endpoint:   conf.OSS.Endpoint,
		bucketName: conf.OSS.BucketName,
	}, nil
}

// UploadFile 上传文件到OSS
func (c *OSSClient) UploadFile(file *multipart.FileHeader, objectKey string) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("打开文件失败: %v", err)
	}
	defer f.Close()

	err = c.bucket.PutObject(objectKey, f)
	if err != nil {
		return "", fmt.Errorf("上传文件失败: %v", err)
	}

	// 生成签名URL，有效期设为1小时
	signedURL, err := c.bucket.SignURL(objectKey, oss.HTTPGet, 3600)
	if err != nil {
		return "", fmt.Errorf("生成签名URL失败: %v", err)
	}

	return signedURL, nil
}

// UploadFileBytes 上传字节数据到OSS并返回可访问的URL
func (c *OSSClient) UploadFileBytes(fileBytes []byte, objectKey string, contentType string) (string, error) {
	err := c.bucket.PutObject(objectKey, bytes.NewReader(fileBytes), oss.ContentType(contentType))
	if err != nil {
		return "", fmt.Errorf("上传文件失败: %v", err)
	}

	// 生成签名URL，有效期设为1小时
	signedURL, err := c.bucket.SignURL(objectKey, oss.HTTPGet, 3600)
	if err != nil {
		return "", fmt.Errorf("生成签名URL失败: %v", err)
	}

	return signedURL, nil
}

// DeleteObject 删除OSS中的文件
func (c *OSSClient) DeleteObject(objectKey string) error {
	// 验证objectKey的合法性
	if objectKey == "" {
		return fmt.Errorf("objectKey不能为空")
	}

	// 确保objectKey不包含bucket信息
	if strings.Contains(objectKey, c.bucketName) {
		return fmt.Errorf("objectKey不应包含bucket信息: %s", objectKey)
	}

	// 确保objectKey格式正确
	if !strings.HasPrefix(objectKey, "avatars/") {
		return fmt.Errorf("非法的objectKey格式: %s", objectKey)
	}

	err := c.bucket.DeleteObject(objectKey)
	if err != nil {
		return fmt.Errorf("删除文件失败: %v", err)
	}
	return nil
}

// GetDomain 获取OSS域名
func (c *OSSClient) GetDomain() string {
	return c.endpoint
}
