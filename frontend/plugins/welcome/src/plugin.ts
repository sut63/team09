import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Workhistory  from './components/Workhistory';
import Personalinformation from './components/Personalinformation';
import Login  from './components/Login';
import Home  from './components/Home';
import Table from './components/Table';
import Training from './components/Training';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/user', Personalinformation);
    router.registerRoute('/history', Workhistory);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/home', Home);
    router.registerRoute('/table', Table);
    router.registerRoute('/training', Training);

  },
});