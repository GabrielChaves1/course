export default defineNuxtRouteMiddleware((to) => {
  const user = useState('user')

  if (!user.value && to.meta.requiresAuth) {
    return navigateTo('/login')
  }

  if (user.value && (to.path === '/login' || to.path === '/register')) {
    return navigateTo('/dashboard')
  }

  if (to.path === '/dashboard' && !user.value) {
    return navigateTo('/login')
  }

  if (to.path === "/confirm") {
    const email = to.query.email as string

    if (!email) {
      return navigateTo('/login')
    }
  }
})

