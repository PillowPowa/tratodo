export function genId(prefix: string = "id") {
  return `${prefix}-${Math.random().toString(36).substr(2, 9)}`;
}
