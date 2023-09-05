FROM loads/alpine:3.8

LABEL maintainer="1570693659@qq.com"

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
ENV WORKDIR /app/main

# 添加应用可执行文件，并设置执行权限
ADD ./main   $WORKDIR/main
ADD ./config/config_docker.toml   $WORKDIR/config/config.toml
ADD ./document/auth/model.conf   $WORKDIR/document/auth/model.conf
ADD ./public/excel/   $WORKDIR/public/excel/
RUN chmod +x $WORKDIR/main
RUN chmod +x $WORKDIR/config/config.toml
RUN chmod +x $WORKDIR/document/auth/model.conf

# 添加静态资源文件
#ADD resource $WORKDIR/resource

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main