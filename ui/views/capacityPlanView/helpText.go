// Copyright (c) 2016 ECS Team, Inc. - All Rights Reserved
// https://github.com/ECSTeam/cloudfoundry-top-plugin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package capacityPlanView

const HelpText = `
**Capacity Plan View**

Capacity plan view shows how many containers of various memory sizes
can be deployed to the foundation based on current capacity.  

NOTE: 
This view currently does not take into consideration the stack
(e.g., cflinuxfs2 vs windows2012R2 ).  This will be a future
enhancement.

**Header information:**

 TODO
 
**Cell list stats:**

  IP - IP address of Cloud Foundry cell
  CPU%% - Total CPU percent consumed by all containers on cell
  RCR - Total reporting containers
  CPUS - Number of CPUs in cell VM
  TOT_MEM - Total Memory in cell VM
  FREE_MEM - Free Memory in cell VM available for containers
  C_RSVD_MEM - Total memory reserved by all containers on cell
  C_USD_MEM - Total memory actually in use by all containers
  TOT_DISK - Total Disk space in cell VM
  FREE_DISK - Free Disk space in cell VM available for containers
  C_RSVD_DSK - Total disk reserved by all containers on cell
  C_USD_DSK - Total disk actually in use by all containers
  MAX_CNTR - Max containers cell can handle
  CNTRS - Number of containers running on cell reported by cell
  DNAME - BOSH Deployment name
  JOB_NAME - BOSH Job name
  JOB_IDX - BOSH Job index
  
**Display: **
Press 'd' to select data view.

**Order / Sort display: **
Press 'o' to show the sort order window allowing multi-column
sorting of any column.

**Clear stats: **
Press shift-C to clear the statistics counters.

**Clipboard menu: **
Press 'c' when a row is selected to open the clipboard menu.
This will copy to clipboard a command you can paste in 
terminal window later.

**Pause display update:**
Press 'p' to toggle pause display update.  When display update is
paused top will continue to capture statstics and display updated
values when unpaused.

**Filter display: **
Press 'f' to show the filter window which allows for filtering
which rows should be displayed

**Reload metadata: **
Press 'r' to force a reload of metadata for app/space/org.  The
metadata is loaded at startup and attempts to stay current by
recognizing when specific data needs to be reloaded. However there
can be circumstances were data becomes stale.

**Refresh screen interval: **
Press 's' to set the sleep time between refreshes. Default
is 1 second.  Valid values are 0.1 - 60.  The refresh interval only
effects how often the client screen is refreshed, it has no effect
on frequency the foundation delivers events. Top uses passive
monitoring for stats, a faster refresh interval will not introduce
additonal load on the CF foundation.

**Select application detail: **
Press UP arrow or DOWN arrow to highlight an application row.
Press ENTER to select the highlighted application and show
additional detail.

**Scroll columns into view: **
Press RIGHT or LEFT arrow to scroll the columns into view if the
window is not wide enough to view all columns.  You can also resize
terminal window to show more columns/rows (resize of cmd.exe window
is not supported on windows while top is running).

**Refresh: **
Press SPACE to force an immediate screen refresh.

**Quit: **
Press 'q' to quit application.

**Log Window: **
Press shift-D to open log window.  This shows internal top
logging messages.  This window will open automatically if any error
message is logged (e.g., connection timeouts).
`
