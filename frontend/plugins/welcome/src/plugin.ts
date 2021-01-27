import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Workhistory  from './components/Workhistory';
import Personalinformation from './components/Personalinformation';
import Login  from './components/Login';
import Home  from './components/Home';
import Training from './components/Training';
import Specialdoctor from './components/Specialdoctor';
import Schedule from './components/Schedule';
import Detail from './components/Detail';
import ScheduleTable from './components/ScheduleTable';



export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/user', Personalinformation);
    router.registerRoute('/history', Workhistory);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/home', Home);
    router.registerRoute('/training', Training);
    router.registerRoute('/specialdoctor', Specialdoctor);
    router.registerRoute('/schedule', Schedule);
    router.registerRoute('/detail', Detail);
    router.registerRoute('/scheduletable', ScheduleTable);


  },
});