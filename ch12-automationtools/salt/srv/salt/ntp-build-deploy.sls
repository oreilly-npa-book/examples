generate_config:
  file.managed:
    - name: /home/ntc/ntp_generated.conf
    - source: salt://ntp_template.j2
    - template: jinja
ntp_peers_example:
  netconfig.managed:
    - template_name: /home/ntc/ntp_generated.conf
    - require:
      - file: /home/ntc/ntp_generated.conf
