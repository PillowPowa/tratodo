export function clickOutside(node: HTMLElement, callback: () => any) {
  const handleClick = ({ target }: MouseEvent) => {
    if (target && !node.contains(target as any)) {
      callback();
    }
  };

  document.addEventListener("click", handleClick, true);

  return {
    destroy() {
      document.removeEventListener("click", handleClick, true);
    },
  };
}
