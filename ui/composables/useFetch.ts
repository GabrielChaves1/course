export function useApiFetch<T>(url: string, options = {}) {
  const config = useRuntimeConfig()
  return $fetch<T>(url, {
    baseURL: config.public.apiBaseUrl,
    credentials: 'include',
    ...options,
  })
}
