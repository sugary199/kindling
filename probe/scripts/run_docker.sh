script_dir="$(dirname "$0")"
workspace_root=$(realpath "${script_dir}/../")

# Docker image information.
#docker_image_with_tag="registry.us-west-1.aliyuncs.com/arms-docker-repo/bpf-compiler:kindling"
docker_image_with_tag="registry.us-west-1.aliyuncs.com/arms-docker-repo/bpf-compiler:kindling-without-extra"

configs=(-v "$HOME/.config:/root/.config" \
  -v "$HOME/.ssh:/root/.ssh" \
  -v "$HOME/.kube:/root/.kube" \
  -v "$HOME/.gitconfig:/root/.gitconfig" \
  -v "$HOME/.arcrc:/root/.arcrc")

IFS=' '
# Read the environment variable and set it to an array. This allows
# us to use an array access in the later command.
read -ra RUN_DOCKER_EXTRA_ARGS <<< "${RUN_DOCKER_EXTRA_ARGS}"

exec_cmd=("/usr/bin/bash")
if [ $# -ne 0 ]; then
  exec_cmd=("${exec_cmd[@]}" "-c" "$*")
fi

docker run --rm -it \
  "${configs[@]}" \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v "${workspace_root}/../:/kindling" \
  -w "/kindling/probe" \
  "${RUN_DOCKER_EXTRA_ARGS[@]}" \
  "${docker_image_with_tag}" \
  "${exec_cmd[@]}"
