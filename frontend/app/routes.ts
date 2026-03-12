import { type RouteConfig, index, route, layout } from '@react-router/dev/routes'

export default [
  index('pages/login/index.tsx'),
  layout('pages/home/index.tsx', [
    route('/home/dashboard', 'pages/dashboard/index.tsx'),
    route('/home/change-password', 'pages/system-management/change-password/index.tsx'),
    route('/home/system-settings', 'pages/system-management/system-settings/index.tsx'),
  ]),
] satisfies RouteConfig
