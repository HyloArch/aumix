export default function useSearchParams() {
  const url = new URL(window.location.href);

  const setSearchParams = (name: string, value: string) => {
    url.searchParams.set(name, value);
    window.history.replaceState({}, "", url);
  };

  const searchParams = {
    get(value: string) {
      return url.searchParams.get(value);
    },
  };

  return {
    searchParams,
    setSearchParams,
  };
}
