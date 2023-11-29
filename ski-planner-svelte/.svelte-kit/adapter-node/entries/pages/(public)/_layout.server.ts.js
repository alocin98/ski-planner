function load({ cookies }) {
  const language = cookies.get("LANG") ?? "de";
  return { language };
}
export {
  load
};
