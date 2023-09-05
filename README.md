# MagnetiQ <img src="./cli/magnetiq.png" width="48" />
Simple tool for adding torrents to the qBittorrent using WebUI using magnet:// protocol. It is convenient to use e.g. in cases where qBittorrent is not running on the local computer, but the user has access to the remote WebUI. Using MagnetiQ, torrents can then be easily added with a single click.


## Installation

Magnetiq contains everything you need in a single binary file and comes in Windows, Linux and Mac versions. On Windows, MagnetiQ can be registered to the magnet:// protocol using the `reg-files` command, which creates `.reg` files for registering and unregistering the protocol. On Linux and Mac, external tools must be used to register the magnet:// protocol.

## Getting started

Before adding torrents to the qBittorrend WebUI using MagnetiQ, it is necessary to configure the connection using a configuration file. Default name is `magnetiq.json` which is located in the same directory as MagnetiQ, however this can be changed using the `-c` parameter. For simplicity, MagnetiQ includes an `init-config` command that creates a ready-to-populate configuration file.   

Once the configuration file is set up correctly, torrents can be added using the `add` command and the `-l` parameter, which specifies the magnet URI to add. This command is also used when MagnetiQ is associated with the magnet:// protocol.

## License

Datel is licensed under the [MIT License](./LICENSE). You are free to use, modify, and distribute this tool in accordance with the terms of the license.

## Disclaimer

Datel is provided "as is" without any warranties, express or implied. The use of this software is at your own risk. The authors and contributors of this project disclaim any and all warranties, including but not limited to, the implied warranties of merchantability and fitness for a particular purpose.

In no event shall the authors or contributors be liable for any direct, indirect, incidental, special, exemplary, or consequential damages (including, but not limited to, procurement of substitute goods or services; loss of use, data, or profits; or business interruption) however caused and on any theory of liability, whether in contract, strict liability, or tort (including negligence or otherwise) arising in any way out of the use of this software, even if advised of the possibility of such damage.