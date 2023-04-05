from nornir import InitNornir
from nornir.core.task import Result, Task
from nornir_jinja2.plugins.tasks import template_string
from nornir_napalm.plugins.tasks import napalm_configure
from nornir_utils.plugins.functions import print_result

# TEMPLATE represents an option to manage multiple templates per platform
TEMPLATE = {
    "eos": "ntp server {{ host['ntp_server'] }}",
    "ios": "ntp server {{ host['ntp_server'] }}",
    "nxos": "ntp server {{  host['ntp_server'] }}",
    "junos": "set system ntp server {{  host['ntp_server'] }}",
}


def config_task(task: Task, template) -> Result:
    """Nornir task that combines two subtasks:
    - Render a configuration from a Jinja2 template
    - Push the rendered configuration to the device
    """
    render_result = task.run(
        task=template_string,
        # The right template per platform is selected
        template=template[task.host.platform],
    )

    config_result = task.run(
        task=napalm_configure,
        # The rendered configuration from previous subtask is used
        # as the configuration input
        configuration=render_result.result,
        # dry_run means the changes without applying them
        dry_run=True,
    )

    return Result(host=task.host, result=config_result)


# Initialize Nornir inventory from a file
nr = InitNornir(config_file="config.yaml")
# The `config_task` will aggregate two subtasks
result = nr.run(
    task=config_task,
    template=TEMPLATE,
)

print_result(result)
