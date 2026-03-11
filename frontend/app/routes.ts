import { type RouteConfig, index, route, layout } from '@react-router/dev/routes'

export default [
  index('pages/login/index.tsx'),
  layout('pages/home/index.tsx', [route('/home/dashboard', 'pages/dashboard/index.tsx')]),
] satisfies RouteConfig
