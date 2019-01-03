// Libraries
import React, {PureComponent, ChangeEvent} from 'react'

// Components
import ProfilePageHeader from 'src/shared/components/profile_page/ProfilePageHeader'
import {Input, IconFont, ComponentSize, EmptyState} from 'src/clockface'
import TaskList from 'src/organizations/components/TaskList'
import FilterList from 'src/shared/components/Filter'

// Types
import {Task} from 'src/api'

interface Props {
  tasks: Task[]
}

interface State {
  searchTerm: string
}

export default class Tasks extends PureComponent<Props, State> {
  constructor(props) {
    super(props)
    this.state = {
      searchTerm: '',
    }
  }

  public render() {
    const {searchTerm} = this.state

    return (
      <>
        <ProfilePageHeader>
          <Input
            icon={IconFont.Search}
            placeholder="Filter tasks..."
            widthPixels={290}
            value={searchTerm}
            onChange={this.handleFilterChange}
            onBlur={this.handleFilterBlur}
          />
        </ProfilePageHeader>
        <FilterList<Task>
          searchTerm={searchTerm}
          searchKeys={['name', 'owner.name']}
          list={this.props.tasks}
        >
          {ts => <TaskList tasks={ts} emptyState={this.emptyState} />}
        </FilterList>
      </>
    )
  }

  private handleFilterBlur = (e: ChangeEvent<HTMLInputElement>): void => {
    this.setState({searchTerm: e.target.value})
  }

  private handleFilterChange = (e: ChangeEvent<HTMLInputElement>): void => {
    this.setState({searchTerm: e.target.value})
  }

  private get emptyState(): JSX.Element {
    return (
      <EmptyState size={ComponentSize.Medium}>
        <EmptyState.Text text="I see nay a task" />
      </EmptyState>
    )
  }
}
