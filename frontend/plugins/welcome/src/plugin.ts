import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import CreateUser  from './components/Work history';
import Login  from './components/Login';
import Home  from './components/Home';
import Table from './components/Table';
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/history', CreateUser);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/home', Home);
    router.registerRoute('/table', Table);
  },
});