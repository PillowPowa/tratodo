export function clickOutside(
  node: HTMLElement,
  callback: (event: MouseEvent) => any
) {
  const handleClick = (e: MouseEvent) => {
    if (!node.contains(e.target as Node)) {
      callback(e);
    }
  };

  document.addEventListener("click", handleClick, true);

  return {
    destroy() {
      document.removeEventListener("click", handleClick, true);
    },
  };
}
