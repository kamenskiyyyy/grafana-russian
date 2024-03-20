import { css } from '@emotion/css';
import React, { ReactElement } from 'react';
import { DragDropContext, Droppable, DropResult } from 'react-beautiful-dnd';

import { selectors } from '@grafana/e2e-selectors';
import { reportInteraction } from '@grafana/runtime';
import { SceneVariable, SceneVariableState } from '@grafana/scenes';
import { useStyles2, Stack, Button, Alert, TextLink } from '@grafana/ui';
import { EmptyState } from '@grafana/ui/src/components/EmptyState/EmptyState';
import { t, Trans } from 'app/core/internationalization';

import { VariableEditorListRow } from './VariableEditorListRow';

export interface Props {
  variables: Array<SceneVariable<SceneVariableState>>;
  onAdd: () => void;
  onChangeOrder: (fromIndex: number, toIndex: number) => void;
  onDuplicate: (identifier: string) => void;
  onDelete: (identifier: string) => void;
  onEdit: (identifier: string) => void;
}

export function VariableEditorList({
  variables,
  onChangeOrder,
  onDelete,
  onDuplicate,
  onAdd,
  onEdit,
}: Props): ReactElement {
  const styles = useStyles2(getStyles);
  const onDragEnd = (result: DropResult) => {
    if (!result.destination || !result.source) {
      return;
    }
    reportInteraction('Variable drag and drop');
    onChangeOrder(result.source.index, result.destination.index);
  };

  return (
    <div>
      <div>
        {variables.length === 0 && <EmptyVariablesList onAdd={onAdd} />}

        {variables.length > 0 && (
          <Stack direction="column" gap={3}>
            <div className={styles.tableContainer}>
              <table
                className="filter-table filter-table--hover"
                data-testid={selectors.pages.Dashboard.Settings.Variables.List.table}
                role="grid"
              >
                <thead>
                  <tr>
                    <th>Variable</th>
                    <th>Definition</th>
                    <th colSpan={5} />
                  </tr>
                </thead>
                <DragDropContext onDragEnd={onDragEnd}>
                  <Droppable droppableId="variables-list" direction="vertical">
                    {(provided) => (
                      <tbody ref={provided.innerRef} {...provided.droppableProps}>
                        {variables.map((variableScene, index) => {
                          const variableState = variableScene.state;
                          return (
                            <VariableEditorListRow
                              index={index}
                              key={`${variableState.name}-${index}`}
                              variable={variableScene}
                              onDelete={onDelete}
                              onDuplicate={onDuplicate}
                              onEdit={onEdit}
                            />
                          );
                        })}
                        {provided.placeholder}
                      </tbody>
                    )}
                  </Droppable>
                </DragDropContext>
              </table>
            </div>
            <Stack>
              <Button
                data-testid={selectors.pages.Dashboard.Settings.Variables.List.newButton}
                onClick={onAdd}
                icon="plus"
              >
                New variable
              </Button>
            </Stack>
          </Stack>
        )}
      </div>
    </div>
  );
}

function EmptyVariablesList({ onAdd }: { onAdd: () => void }): ReactElement {
  return (
    <EmptyState
      buttonLabel={t('variables.empty-state.button-title', 'Add variable')}
      message={t('variables.empty-state.title', 'There are no variables added yet')}
      onButtonClick={onAdd}
    >
      <Alert severity="info" title={t('variables.empty-state.info-box-title', 'What do variables do?')}>
        <p>
          <Trans i18nKey="variables.empty-state.info-box-content">
            Variables enable more interactive and dynamic dashboards. Instead of hard-coding things like server or
            sensor names in your metric queries you can use variables in their place. Variables are shown as list boxes
            at the top of the dashboard. These drop-down lists make it easy to change the data being displayed in your
            dashboard.
          </Trans>
        </p>
        <Trans i18nKey="variables.empty-state.info-box-content-2">
          Check out the{' '}
          <TextLink external href="https://grafana.com/docs/grafana/latest/variables/">
            Templates and variables documentation
          </TextLink>{' '}
          for more information.
        </Trans>
      </Alert>
    </EmptyState>
  );
}

const getStyles = () => ({
  tableContainer: css({
    overflow: 'scroll',
    width: '100%',
  }),
});
