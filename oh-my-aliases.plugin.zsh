__oh_my_aliases__PLUGIN_DIR=${0:a:h}

__oh_my_aliases__DELIMITER="__oh_my_aliases__DELIMITER"

_oh_my_aliases__fn () {
  local alias_list
  alias_list=$(alias)

  local history
  history=$(history -200)

  echo "$alias_list\n__oh_my_aliases__DELIMITER\n$history" | python3 ${__oh_my_aliases__PLUGIN_DIR}/oh-my-aliases.py $*
}

autoload -Uz add-zsh-hook
add-zsh-hook preexec _oh_my_aliases__fn
