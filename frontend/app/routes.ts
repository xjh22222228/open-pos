import { type RouteConfig, index, route, layout } from '@react-router/dev/routes'

export default [
  index('pages/login/index.tsx'),
  layout('pages/home/index.tsx', [
    route('/home/dashboard', 'pages/dashboard/index.tsx'),
    route('/home/change-password', 'pages/system-mgr/change-password/index.tsx'),
    route('/home/system-settings', 'pages/system-mgr/system-settings/index.tsx'),
    route('/home/personal-Info', 'pages/system-mgr/personal-Info/index.tsx'),
  ]),
] satisfies RouteConfig
