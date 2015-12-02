
/**
 * Renders a list of service plans
 */

import React from 'react';
import Reactable from 'reactable';

import Button from './button.jsx';
import serviceActions from '../actions/service_actions.js';
import ServicePlanStore from '../stores/service_plan_store.js';

var Table = Reactable.Table,
    Thead = Reactable.Thead,
    Th = Reactable.Th,
    Tr = Reactable.Tr,
    Td = Reactable.Td;

export default class ServicePlanList extends React.Component {
  constructor(props) {
    super(props);
    this.props = props;
    this.state = {
      serviceGuid: props.initialServiceGuid,
      servicePlans: this.props.initialServicePlans
    };
  }

  componentWillReceiveProps(nextProps) {
    this.setState({
      serviceGuid: nextProps.initialServiceGuid,
      servicePlans: nextProps.initialServicePlans
    });
  }

  _handleAdd = (planGuid) => {
    serviceActions.createInstanceDialog(this.state.serviceGuid,
      planGuid);
  }

  get columns() {
    var columns = [
      { label: 'Name', key: 'label' },
      { label: 'Free', key: 'free' },
      { label: 'Description', key: 'description' },
      { label: 'Cost', key: 'extra.costs.amount.usd' },
      { label: 'Actions', key: 'actions' }
    ];

    return columns;
  }

  get rows() {
    var rows = this.state.servicePlans;
    return rows;
  }

  render() {
    var content = <h4 className="test-none_message">No service plans</h4>;
    if (this.state.servicePlans.length) {
      content = (
      <Table sortable={ true } className="table">
        <Thead>
          { this.columns.map((column) => {
            return (
              <Th column={ column.label } className={ column.key }>
                { column.label }</Th>
            )
          })}
        </Thead>
        { this.state.servicePlans.map((plan) => {
          return (
            <Tr key={ plan.guid }>
              <Td column={ this.columns[0].label }>
                <span>{ plan.name }</span></Td>
              <Td column={ this.columns[1].label }>{ plan.free }</Td>
              <Td column={ this.columns[2].label }>{ plan.updated_at }</Td>
              <Td column={ this.columns[3].label }>
                <span>
                  ${ (plan.extra.costs[0].amount.usd || 0).toFixed(2) } monthly
                </span>
              </Td>
              <Td column={ this.columns[4].label }>
                <Button 
                  onClickHandler={ this._handleAdd.bind(this, plan.guid) } 
                  label="create">
                  <span>Create service instance</span>
                </Button>
              </Td>
            </Tr>
          )
        })}
      </Table>
      );
    }

    return (
    <div className="tableWrapper"> 
      { content }
    </div>
    );
  }
}

ServicePlanList.propTypes = {
  initialServiceGuid: React.PropTypes.string,
  initialServicePlans: React.PropTypes.array
};

ServicePlanList.defaultProps = {
  initialServicePlans: []
};
