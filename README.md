# dockerfile多阶段编译 参考资料

 1. https://medium.com/@tonistiigi/advanced-multi-stage-build-patterns-6f741b852fae
 2. https://docs.docker.com/develop/develop-images/multistage-build/

# docker build 参数
 --target <阶段名称>

通过改变阶段名称,build不同的镜像版本
