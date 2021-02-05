import { createPlugin } from '@backstage/core';
import Workhistory  from './components/Workhistory';
import Personalinformation from './components/Personalinformation';
import Login  from './components/Login';
import Home  from './components/Home';
import Training from './components/Training';
import Specialdoctor from './components/Specialdoctor';
import Schedule from './components/Schedule';
import Detail from './components/Detail';
import WorkhistoryTable from './components/WorkhistoryTable';
import SpecialdoctorTable from './components/SpecialdoctorTable';
import TrainingTables from './components/TrainingTables';
import ScheduleTable from './components/ScheduleTable';
import PersonalinformationTable from './components/PersonalinformationTable';
import DetailTable from './components/DetailTable';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/user', Personalinformation);
    router.registerRoute('/history', Workhistory);
    router.registerRoute('/home', Home);
    router.registerRoute('/training', Training);
    router.registerRoute('/specialdoctor', Specialdoctor);
    router.registerRoute('/schedule', Schedule);
    router.registerRoute('/detail', Detail);
    router.registerRoute('/workhistorytable', WorkhistoryTable);
    router.registerRoute('/specialdoctortable', SpecialdoctorTable);
    router.registerRoute('/trainingtable', TrainingTables);
    router.registerRoute('/scheduletable', ScheduleTable);
    router.registerRoute('/personalinformationtable', PersonalinformationTable);
    router.registerRoute('/detailtable', DetailTable);

  },
});