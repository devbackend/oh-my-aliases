_oh_my_aliases__fn () {
  local alias_list
  alias_list=$(alias)

  echo "$alias_list" | oh-my-aliases $*
}

autoload -Uz add-zsh-hook
add-zsh-hook preexec _oh_my_aliases__fn