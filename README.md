# dockerfile多阶段编译 参考资料

 1. https://medium.com/@tonistiigi/advanced-multi-stage-build-patterns-6f741b852fae
 2. https://docs.docker.com/develop/develop-images/multistage-build/

# Makefile 参数
## make docker 
    @docker build --target $(stage_version) -t goimage:$(tag) . 
    --target <阶段名称>
    通过改变阶段名称,build不同的镜像版本
    eg. make docker tag=2.1 stage_version=stage_with_timezone

    不同版本可以基于同一个阶段镜像,具体用法见参考资料1
