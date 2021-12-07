FROM 949191617935.dkr.ecr.cn-north-1.amazonaws.com.cn/jm-app/jm-app-base:2.0.4
ADD report /report

ARG AE_VERSION
# 添加资源文件，并解压
ADD build/ae/ae_data_v2/$AE_VERSION/lookups-$AE_VERSION.tar.gz /ae_data/lookups
ADD build/ae/ae_data_v2/$AE_VERSION/lua_src-$AE_VERSION.tar.gz /ae_data/lua_src
ADD build/ae/ae_data_v2/$AE_VERSION/question-$AE_VERSION.tar.gz /ae_data/question
ADD build/ae/ae_data_v2/$AE_VERSION/biz_conf-$AE_VERSION.tar.gz /ae_data/biz_conf

ENTRYPOINT ["/report"]
