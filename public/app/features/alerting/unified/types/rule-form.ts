import { AlertQuery, GrafanaAlertStateDecision } from 'app/types/unified-alerting-dto';

import { Folder } from '../components/rule-editor/RuleFolderPicker';

export enum RuleFormType {
  grafana = 'grafana',
  cloudAlerting = 'cloud-alerting',
  cloudRecording = 'cloud-recording',
}

export interface ContactPoint {
  alertManager: string;
  selectedContactPoint: string;
  overrideGrouping: boolean;
  groupBy: string[];
  overrideTimings: boolean;
  groupWaitValue: string;
  groupIntervalValue: string;
  repeatIntervalValue: string;
  muteTimeIntervals: string[];
}

export interface RuleFormValues {
  // common
  name: string;
  type?: RuleFormType;
  dataSourceName: string | null;
  group: string;

  labels: Array<{ key: string; value: string }>;
  annotations: Array<{ key: string; value: string }>;

  // grafana rules
  queries: AlertQuery[];
  condition: string | null; // refId of the query that gets alerted on
  noDataState: GrafanaAlertStateDecision;
  execErrState: GrafanaAlertStateDecision;
  folder: Folder | null;
  evaluateEvery: string;
  evaluateFor: string;
  isPaused?: boolean;
  manualRouting: boolean; // if true contactPoints are used
  contactPoints?: ContactPoint[];

  // cortex / loki rules
  namespace: string;
  forTime: number;
  forTimeUnit: string;
  keepFiringForTime?: number;
  keepFiringForTimeUnit?: string;
  expression: string;
}
